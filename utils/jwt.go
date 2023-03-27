package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/joho/godotenv"

	"mini-project-internship/models/entity"
)

func GenerateTokenJwt(claims *jwt.MapClaims) (string, error) {
	godotenv.Load()
	key := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	webToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return webToken, nil
}

func GenerateClaims(user entity.User) jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["nama"] = user.Nama
	claims["email"] = user.Email
	claims["is_admin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()
	return claims
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	godotenv.Load()
	key := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(key), nil
	})
	
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}