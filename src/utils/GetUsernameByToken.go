package utils

import (
	"net/http"

	"github.com/taylorwebk/kafei-api/src/structs"

	"github.com/dgrijalva/jwt-go"
	"github.com/taylorwebk/kafei-api/src/config"
)

// GetUsernameByToken returns a username
func GetUsernameByToken(tokenstr string, w http.ResponseWriter) (string, uint) {
	token, err := jwt.ParseWithClaims(tokenstr, &userClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.KeyJWT), nil
	})
	if err != nil {
		JSONResponse(http.StatusInternalServerError, structs.Response{Message: err.Error()}, w)
	}
	if claims, ok := token.Claims.(*userClaim); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		return claims.Username, claims.ID
	}
	return "", 0
}
