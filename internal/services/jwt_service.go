package services

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secretKey []byte
}

func (service *JwtService) GenerateToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(service.secretKey)

	if err != nil {
		log.Fatalf("Could not generate token: %v", err)
	}

	return tokenString
}

func (service *JwtService) ValidateToken(token string, claims jwt.MapClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return service.secretKey, nil
	})
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
}
