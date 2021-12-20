package users

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt"
)

//TODO use env var?
const hmacSampleSecret = "SRGNKNQJBKBQKJSBKJBKJQBSKJQKJSDBKJBKJ"

func generateJwt(user *User) (string, error) {
	fmt.Println("generating jwt for %v", user.Username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"username": user.Username,
		"created": user.Created,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	return tokenString, err;
}

func validateJwt(tokenString string) (*User, error) {
	fmt.Println("Validating jwt %v", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSampleSecret), nil
	})
	
	fmt.Println("will claims jwt")
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println(claims["username"], claims["userId"], claims["created"])
		t1, timeErr := time.Parse(time.RFC3339, claims["created"].(string))
		if timeErr != nil {
			fmt.Errorf("Error parsin time from jwt", timeErr)
		}
		id := claims["userId"].(float64)
			return &User {
				ID:			int64(id),
				Username: 	claims["username"].(string),
				Created: 	t1,
			}, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}