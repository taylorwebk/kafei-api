package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// EndActivity ends an activity
func EndActivity(w http.ResponseWriter, r *http.Request) {
	var err error
	db := database.SQL
	acr := act{}
	usertoken := r.Context().Value("user").(structs.UserToken)
	err = json.NewDecoder(r.Body).Decode(&acr)
	if err != nil {
		utils.JSONResponse(
			http.StatusInternalServerError,
			err.Error(),
			w,
		)
		return
	}
	activity := structs.Activity{}
	db.Where("id = ?", acr.ID).First(&activity)
	activity.Status = "ENDED"
	db.Save(&activity)
	response := structs.Response{
		Message: fmt.Sprintf("Finalizo la actividad: %s.", activity.Literal),
		Content: structs.Token{
			Token: utils.GenerateToken(usertoken.ID, usertoken.Username, w),
		},
	}
	utils.JSONResponse(
		http.StatusOK,
		response,
		w,
	)
}

type act struct {
	ID uint `json:"activity_id"`
}
