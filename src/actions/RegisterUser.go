package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
	"golang.org/x/crypto/bcrypt"
)

// Hello hola mundo
func Hello(w http.ResponseWriter, r *http.Request) {
	text, _ := json.Marshal("Welcome to the rest api of KAFEI pro by taylorwebk")
	w.Write(text)
	fmt.Print("hola")
}

// RegisterUser register a User in the database
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)
	db := database.SQL
	user := &structs.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	var query []structs.User
	db.Where("username = ? OR email = ?", user.Username, user.Email).Find(&query)
	if len(query) > 0 {
		res := structs.Response{
			Message: fmt.Sprintf(
				"Ups, el username: %s o email: %s ya esta siendo usado.",
				user.Username,
				user.Email,
			),
		}
		utils.JSONResponse(http.StatusConflict, res, w)
	} else {
		var hashedPassword []byte
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		user.Password = string(hashedPassword)
		db.Create(&user)
		res := structs.Response{
			Message: "Todo listo para ser más efectivo y comenzar a ahorrar.",
			Content: structs.Token{
				Token: utils.GenerateToken(user.ID, user.Username, w),
				Data:  user.NewStruct(),
			},
		}
		utils.JSONResponse(http.StatusCreated, res, w)
	}
}
