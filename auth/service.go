package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type JwtService struct {
}

var SECRET_KEY = []byte("BWASTARTUP_s3cr3T_K3Y")

func NewService() *JwtService {
	return &JwtService{}
}

func (s *JwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["User_Id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	SignedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return SignedToken, err
	}

	return SignedToken, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
