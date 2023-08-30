package restapi

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var authenticated bool

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		err := fmt.Errorf(err.Error())
		authenticated = false // Set authentication status to false
		return nil, err
	}
	if token.Valid {
		authenticated = true // Set authentication status to true
		return claims["user"].(string), nil
	}
	authenticated = false // Set authentication status to false
	return nil, errors.New("invalid token")
}
