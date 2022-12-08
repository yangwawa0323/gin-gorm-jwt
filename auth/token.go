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
)

var debug = utils.Debug
var Errors = utils.Errors

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

// TODO: The secret key should put in conf/app.yaml
var jwtkey = []byte("superscretkey")

// GenerateToken: generate a formatted token string for a models.User
func GenerateToken(user *models.User) (tokenString string, err error) {
	expiresAt := time.Now().Add(3 * time.Hour)
	claims := &JWTClaim{
		Email:    user.Email,
		Username: user.Username,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtkey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		// func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(jwtkey), nil
		// },
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtkey, nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims) //
	if !ok {
		err = errors.New(Errors[utils.ParseClaimError])
		return
	}
	if claims["exp"].(int64) < time.Now().Local().Unix() {
		err = errors.New(Errors[utils.TokenHasExpired])
		return
	}
	return
}

// ExtractTokenString: extract token string from user request context
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
func ExtractTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			// }
			return jwtkey, nil
		})
	if err == nil && token.Valid {
		return token, nil
	}
	return nil, errors.New(Errors[utils.TokenIsInvalid])
}

// Extract the custom claim from token string
// IMPORTANT: custom claim to json is ok, but the field name has been changed.
// We cannot read the struct from token.Claims by type assetion, it only `nil` we got.
// To resolve the problem , we has use jwt.MapClaims, by access the key/value.
func ExtractTokenClaim(tokenString string) (jwt.MapClaims, error) {
	if err := ValidateToken(tokenString); err != nil {
		return nil, err
	}
	token, err := ExtractTokenFromString(tokenString)
	if err != nil {
		return nil, err
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	// debug(fmt.Sprintf("%#v", token.Claims))
	if !ok {
		return nil, errors.New(Errors[utils.ParseClaimError])
	}
	return claim, err
}
