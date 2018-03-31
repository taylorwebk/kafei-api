package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/taylorwebk/kafei-api/src/config"
)

// GenerateToken generate token based in user's username
func GenerateToken(id uint, username string, w http.ResponseWriter) string {
	claims := userClaim{
		id,
		username,
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
	}
	tokenInstance := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenInstance.SignedString([]byte(config.Conf.KeyJWT))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "null"
	}
	return token
}

type userClaim struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
