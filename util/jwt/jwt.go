package jwt

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string) (*string, error) {
	claims := jwt.MapClaims{}
	claims["user_name"] = username
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := newToken.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))

	return &token, err
}

func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}

	bearerToken := ctx.GetHeader("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return bearerToken
}

func validateTokenMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(os.Getenv("JWT_SIGNATURE")), nil
}

func TokenValid(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)

	_, err := jwt.Parse(tokenString, validateTokenMethod)
	if err != nil {
		return err
	}

	return nil
}
