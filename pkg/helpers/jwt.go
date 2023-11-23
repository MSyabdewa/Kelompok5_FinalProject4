package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/infra/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id int, email string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
	})

	tokenString, err := token.SignedString([]byte(config.GetAppConfig().JWTSecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("a token is required to access")
	getToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(getToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	tokenString := strings.Split(getToken, " ")[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.GetAppConfig().JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
