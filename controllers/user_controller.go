/**
*   IMPORTANT!!!
*   You can generate token from customized claim, but you `CANNOT` get the struct from token string,
*   because the **jwt.NewWithClaims** func that accordings the json tag to generate the token string
*   that cause all the field name is diffrence against the custom claim struct's.
*   There is a way to get data by type assetion to **jwt.MapClaim**, and use key/value to get the data.
 */

package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/auth"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/services"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gorm.io/gorm"
)

var debug = utils.Debug
var errorDebug = utils.ErrorDebug

// RegisterUser func has three steps
// 1. save the user to DB
// 2. send a activate mail
// 3. generate a JWT token
// finally. return the JSON response
// Testing at 2022-12-08 12:45AM
func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	// Save new user
	usersvc := services.NewUserService(&user)
	user.UserClass = models.Guest // new user by default is a **guest**
	result := usersvc.New(&user)

	if result != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error(),
		})
		ctx.Abort()
		return
	}

	var sendMail = make(chan error, 1)
	var generateJWT = make(chan error, 1)
	// send a activate mail
	go func() {
		debug("Next step: send a activate mail.")
		body, err := user.GenerateActivateMailBody()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		mailDialer := models.NewMailDialer(
			"Welcome to register 51cloudclass website",
			body,
			user,
		)
		sendMail <- mailDialer.SendMail_gomailV2()
	}()

	go func() {
		debug("Finally generate JWT")
		token, err := auth.GenerateToken(&user)
		generateJWT <- err
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"userId":   user.ID,
			"email":    user.Email,
			"username": user.Username,
			"token":    token,
		})
	}()

	if <-sendMail != nil {
		debug("failed to send email")
	}
	if <-generateJWT != nil {
		debug("failed to generate JWT")
	}
}

// TODO: testing
func ConfirmMailActivate(ctx *gin.Context) {
	type Secret struct {
		Token string `form:"token"`
	}
	var svcUser = models.User{}
	var mailSecret Secret
	userIDparam, ok := ctx.Params.Get("user_id")

	if err := ctx.ShouldBindQuery(&mailSecret); err != nil {
		errorDebug(err)
	}
	userID, err := strconv.Atoi(userIDparam)

	if ok && err == nil {
		usersvc := services.NewUserService(&svcUser)
		user, dbErr := usersvc.FindUserByID(int64(userID))
		unescapedSecret, tokenErr := url.QueryUnescape(mailSecret.Token)

		if dbErr == nil && tokenErr == nil &&
			user.IsActivateMailStringValid([]byte(unescapedSecret)) {
			ctx.Redirect(http.StatusTemporaryRedirect, "/")
		} else {
			debug("[DEBUG]: activate token is invalid")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": Errors[utils.TokenHasExpired],
			})
		}
	}
}

// TODO: not implemented yet
func UploadAvator(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "upload avator",
	})
}

// TODO: not implemented yet
func ChangePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "change passwod",
	})
}

// TODO: check JWT token valid
// 1. User never login or token has expired
// 2. User has token and valid.
func Login(ctx *gin.Context) {

	// Method 1. authenticate user by token
	tokenString := auth.ExtractTokenString(ctx)
	claim, err := auth.ExtractTokenClaim(tokenString) // It returns jwt.MapClaim
	debug(fmt.Sprintf("%#v", claim))
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"username": claim["username"],
			"email":    claim["email"],
			"user_id":  claim["user_id"],
			"expireAt": claim["exp"],
		})
	} else {
		// Method 2. authenticate user by form data.
		var user models.User = models.User{}
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, &user)
			return
		} else {
			usersvc := services.NewUserService(&user)
			result := usersvc.DB.Where("name = ? and password = ?",
				usersvc.User.Name,
				usersvc.User.Password).First(&user)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": Errors[utils.UserNotExists],
				})
				return
			}
		}
		// TODO: GernateToken for this user.
		tokenString, err := auth.GenerateToken(&user)
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "user has logged in",
			"token":   tokenString,
		})
	}
}

// TODO: not implemented yet
func ListMessages(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "list messages",
	})
}

// TODO: not implemented yet
func GetUserInfoFromCookie(ctx *gin.Context) *models.User {
	// ctx.Cookie()
	debug("Not implemented yet")
	return nil
}