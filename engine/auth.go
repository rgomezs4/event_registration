package engine

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJwt generates a new jwt token based on the user_id
func GenerateJwt(userID int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 10).Unix(),
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

// VerifyJwt verifies that the jwt provided matches the criteria for the request
func VerifyJwt(userID int, token string) map[string]interface{} {
	response := make(map[string]interface{})
	response["is_valid"] = false

	jwtToken, err := decodeToken(token)
	if err != nil {
		response["message"] = err.Error()
		return response
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok && !jwtToken.Valid {
		// If the token isn't valid it returns an "Invalid token" message
		response["message"] = "Invalid token"
		return response
	}

	if err := claims.Valid(); err != nil {
		response["message"] = err.Error()
		return response
	}

	if claims["user_id"].(float64) != float64(userID) {
		response["message"] = "invalid token"
		return response
	}

	response["is_valid"] = true
	return response
}

func decodeToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unable to decode the token")
		}
		return []byte(os.Getenv("APP_KEY")), nil
	})
}
