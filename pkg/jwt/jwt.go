package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Secret = []byte("secret")
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	// Create a map to store our claim
	claims := token.Claims.(jwt.MapClaims)

	// Set token claims
	claims["username"] = username
	// TODO ADD USERID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses a jwt and returns the username
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", err
}
