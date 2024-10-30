package utility

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte("oklajsjdmmasdkkasdkkllasodo")
)

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 60).Unix(),
			"type":     "primary-token",
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func RefreshToken(username string) (string, any, error) {
	Expires := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      Expires,
			"type":     "refresh-token",
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", "", nil
	}
	return tokenString, Expires, nil
}

func CheckBearer(header string) (string, error) {

	if header == "" {
		return "", errors.New("Header can't blank.")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil

}

func VerifyToken() func(c *gin.Context) {

	return func(c *gin.Context) {

		tokenString, err := CheckBearer(c.GetHeader("Authorization"))
		if err != nil {
			ResponseErrorCustom(c, http.StatusBadRequest, nil, err.Error())
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			fmt.Println("parse : ", token)
			return secretKey, nil
		})

		if err != nil {
			fmt.Println(err)
			ResponseErrorCustom(c, http.StatusUnauthorized, nil, err.Error())
			c.Abort()
			return
		}

		if !token.Valid {
			ResponseErrorCustom(c, http.StatusBadRequest, nil, "Token not valid.")
			c.Abort()
			return
		}

		c.Next()

	}

}

func ValidToken(c *gin.Context, token string) bool {

	// Parse
	tokens, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false
	}

	if !tokens.Valid {
		return false
	}

	return true

}
