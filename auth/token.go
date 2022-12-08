package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	myerr "github.com/yangwawa0323/gin-gorm-jwt/utils/errors"
)

var debug = utils.Debug
var Errors = myerr.Errors

type JWTClaim struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	UserID    uint   `json:"user_id"`
	AvatarURL string `json:"avatar_url"`
	jwt.StandardClaims
}

// TODO: The secret key should put in conf/app.yaml
var jwtkey = []byte("superscretkey")

// GenerateToken: generate a formatted token string for a models.User
func GenerateToken(user *models.User) (tokenString string, err error) {
	expiresAt := time.Now().Add(3 * time.Hour)
	claims := &JWTClaim{
		Email:     user.Email,
		Username:  user.Username,
		UserID:    user.ID,
		AvatarURL: user.AvatarURL,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtkey)
	return
}

// func ValidateToken(signedToken string) bool {
// 	claim, err := ParseClaim(signedToken)
// 	debug(signedToken)
// 	debug(fmt.Sprintf("%#v", claim))
// 	return err == nil
// }

func ParseClaim(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		// Callback func to get jwt secret key
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%s %v", Errors[myerr.UnexpectedSigningMethod],
					token.Header["alg"])
			}
			return jwtkey, nil
		},
	)
	if err != nil {
		return
	}
	var ok bool
	claims, ok = token.Claims.(*JWTClaim) //
	// debug(fmt.Sprintf("%#v: %#v", token.Claims, claims))
	if !ok {
		err = errors.New(Errors[myerr.ParseClaimError])
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New(Errors[myerr.TokenHasExpired])
		return
	}
	return
}

// ExtractTokenString: extract token signed string from user request context
func ExtractTokenString(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		if strings.Contains(bearerToken, "Bearer ") {
			return strings.Split(bearerToken, "Bearer ")[1]
		}
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenFromString: generate token from given token string
func ExtractTokenFromString(signedString string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedString,
		func(token *jwt.Token) (interface{}, error) {
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			// }
			return jwtkey, nil
		})
	if err == nil && token.Valid {
		return token, nil
	}
	return nil, errors.New(Errors[myerr.TokenIsInvalid])
}

// Extract the custom claim from token string
// IMPORTANT: custom claim to json is ok, but the field name has been changed.
// We cannot read the struct from token.Claims by type assetion, it only `nil` we got.
// To resolve the problem , we has use jwt.MapClaims, by access the key/value.
