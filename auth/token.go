package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
)

var debug = utils.Debug

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserID   uint   `json:"user_id"`
	jwt.StandardClaims
}

var jwtkey = []byte("superscretkey")

func GenerateJWT(user *models.User) (tokenString string, err error) {
	expirationTime := time.Now().Add(3 * time.Hour)
	claims := &JWTClaim{
		Email:    user.Email,
		Username: user.Username,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
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
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtkey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ExtractToken(ctx *gin.Context) string {

	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, ".")) == 2 {
		return strings.Split(bearerToken, ".")[1]
	}
	debug("Not test yet")
	return ""
}

func ExtractTokenUserID(ctx *gin.Context) (uint, error) {
	tokenString := ExtractToken(ctx)
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtkey, nil
		})
	if err != nil {
		return 0, err
	}
	claim, ok := token.Claims.(JWTClaim)
	if ok && token.Valid {

		uid, err := strconv.ParseUint(fmt.Sprintf("%d", claim.UserID), 10, 32)
		if err != nil {
			debug("cannot parse user id in JWTClaim")
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
