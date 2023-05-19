package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	UserID  int 		
	Expired time.Time
}

const (
	TOKEN_KEY = "U64w7c3xpx6v5kzM"
	TOKEN_EXPIRED = 10 * 60 * time.Second
)

func GenerateToken(userId int) (string, error) {
	payload := Payload{
		UserID: userId,
		Expired: time.Now().Add(TOKEN_EXPIRED),
	}

	claims := jwt.MapClaims{
		"payload": payload,
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString([]byte(TOKEN_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(TOKEN_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		payloadInterface := claims["payload"]

		payload := Payload{}

		payloadByte, err := json.Marshal(payloadInterface)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(payloadByte, &payload)

		if err != nil {
			return nil, err
		}

		now := time.Now()

		if now.After(payload.Expired) {
			return nil, errors.New("Token has been expired.")
		}

		return &payload, nil
	}else {
		return nil, err
	}
}