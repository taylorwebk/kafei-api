package actions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"

	"github.com/taylorwebk/kafei-api/src/utils"
)

// AddInterval ads an interval to activity
func AddInterval(w http.ResponseWriter, r *http.Request) {
	var err error
	db := database.SQL
	intStruct := intervalStruct{}
	usertoken := r.Context().Value("user").(structs.UserToken)
	err = json.NewDecoder(r.Body).Decode(&intStruct)
	if err != nil {
		utils.JSONResponse(
			http.StatusInternalServerError,
			err.Error(),
			w,
		)
		return
	}
	start, err := time.Parse("2006-01-02 15:04:05", intStruct.StartedAt)
	if err != nil {
		utils.JSONResponse(
			http.StatusBadRequest,
			"Campo de fecha no aceptado",
			w,
		)
	}
	end := start.Add(time.Second * time.Duration(intStruct.Seconds))
	activity := structs.Activity{}
	db.Where("id = ?", intStruct.ActivityID).First(&activity)
	interval := structs.Interv{
		Start: start,
		End:   end,
	}
	activity.AddInterval(interval)
	db.Save(&activity)
	tokenstr := utils.GenerateToken(usertoken.ID, usertoken.Username, w)
	response := structs.Response{
		Message: fmt.Sprintf("Intervalo de: %s guardado.", activity.Literal),
		Content: structs.Token{
			Token: tokenstr,
		},
	}
	utils.JSONResponse(
		http.StatusOK,
		response,
		w,
	)
}

type intervalStruct struct {
	ActivityID uint   `json:"activity_id"`
	StartedAt  string `json:"startedat"`
	Seconds    uint64 `json:"seconds"`
}
