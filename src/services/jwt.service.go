package services

import (
	"fmt"
	"restapi/src/configs"
	"restapi/src/helpers"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTServivce struct {
	secretKey string
	issuer    string
}

type jwtCustomClaim struct {
	UserID string `json:"uid"`
	jwt.StandardClaims
}

// NewJWTService is function
func NewJWTService() *JWTServivce {
	return &JWTServivce{
		secretKey: helpers.GetenvWithDefaultValue("JWT_SECRET_KEY", "secret_key"),
		issuer:    configs.Issuer,
	}
}

// GenerateToken is function
func (j *JWTServivce) GenerateToken(UserID string) string {
	startTime := time.Now()

	fmt.Printf("Issuer: %s", j.issuer)

	// tao nhung thanh phan kinh kem vao token
	claims := &jwtCustomClaim{
		UserID,
		jwt.StandardClaims{
			Id:        UserID,
			Subject:   "API Authorization",
			Audience:  "ducthuan.net",
			Issuer:    j.issuer,
			IssuedAt:  startTime.Unix(),
			ExpiresAt: startTime.Add(configs.ExpiredTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(configs.SigningMethod, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return t
}

// ValidateToken is function
func (j *JWTServivce) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
