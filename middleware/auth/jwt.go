package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/saisesai/www.sxu.ink/model"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Abort()
			ctx.String(http.StatusOK, "请先登录")
			return
		}

		claims, err := ParseToken(auth)
		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				newToken, _ := reNewToken(claims)
				if newToken != "" {
					ctx.Header("newToken", newToken)
					ctx.Request.Header.Set("Authorization", newToken)
					ctx.Next()
					return
				}
			}
			ctx.Abort()
			ctx.String(http.StatusOK, err.Error())
			return
		}
		ctx.Next()
	}
}

type Claims struct {
	User model.User
	jwt.StandardClaims
}

const (
	TokenExpireDuration = time.Hour * 2
)

var Secret = []byte("secret")

func GenerateToken(userInfo model.User) (string, error) {
	var (
		expirationTime = time.Now().Add(TokenExpireDuration)
		claims         = &Claims{
			User: userInfo,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
				Issuer:    "name",
			},
		}
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	var (
		claims = &Claims{}
	)

	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	return claims, err
}

func reNewToken(claims *Claims) (string, error) {
	if withinLimit(claims.ExpiresAt, 600) {
		return GenerateToken(claims.User)
	}

	return "", errors.New("登录已过期, 请重新登录")
}

func withinLimit(s int64, l int64) bool {
	e := time.Now().Unix()
	return e-s < l
}
