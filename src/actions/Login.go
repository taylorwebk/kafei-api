package actions

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// Login handles the login of an user
func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	userCredentials := &loginCredentials{}
	db := database.SQL
	user := structs.User{}
	err = json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		panic(err)
	}
	db.Where(&structs.User{
		Username: userCredentials.Username,
	}).First(&user)
	if user.ID == 0 {
		db.Where(&structs.User{
			Email: userCredentials.Username,
		}).First(&user)
	}
	if user.ID == 0 {
		userOrPasswordInvalid(w)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password))
	if err != nil {
		userOrPasswordInvalid(w)
		return
	}
	response := structs.Response{
		Message: "Es grato tenerte de vuelta.",
		Content: structs.Token{
			Token: utils.GenerateToken(user.ID, user.Username, w),
			Data:  user.NewStruct(),
		},
	}
	utils.JSONResponse(http.StatusOK, response, w)
}
func userOrPasswordInvalid(w http.ResponseWriter) {
	utils.JSONResponse(
		http.StatusUnauthorized,
		structs.Response{
			Message: "Datos de ingreso no válidos.",
		},
		w,
	)
}

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
