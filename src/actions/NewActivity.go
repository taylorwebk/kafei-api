package actions

import (
	"encoding/json"
	"net/http"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// NewActivity register a new activity for the user
func NewActivity(w http.ResponseWriter, r *http.Request) {
	db := database.SQL
	var err error
	activity := structs.Activity{}
	user := structs.User{}
	err = json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		utils.JSONResponse(
			http.StatusUnprocessableEntity,
			structs.Response{
				Message: err.Error(),
			},
			w,
		)
		return
	}
	usertoken := r.Context().Value("user").(structs.UserToken)
	db.Where("id = ?", usertoken.ID).First(&user)
	user.AddActivity(activity)
	db.Save(&user)
	res := getActivities(user.ID)
	tokenstr := utils.GenerateToken(user.ID, user.Username, w)
	response := structs.Response{
		Message: "Nueva actividad guardada.",
		Content: structs.Token{
			Token: tokenstr,
			Data:  res,
		},
	}
	utils.JSONResponse(http.StatusOK, response, w)
}
