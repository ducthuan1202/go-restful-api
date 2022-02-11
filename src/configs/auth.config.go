package configs

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	ExpiredTime   = time.Hour * 24 * 30 // 30 ngay
	SigningMethod = jwt.SigningMethodHS256
	Issuer        = "n-d-t-12-02"
)
