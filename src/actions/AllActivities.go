package actions

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// AllActivities get all activities from a user
func AllActivities(w http.ResponseWriter, r *http.Request) {
	usertoken := r.Context().Value("user").(structs.UserToken)
	//spew.Dump(getActivities(usertoken.ID))
	responseToken := structs.Token{
		Token: utils.GenerateToken(usertoken.ID, usertoken.Username, w),
		Data:  getActivities(usertoken.ID),
	}
	response := structs.Response{
		Message: "Actividades cargadas correctamente.",
		Content: responseToken,
	}
	utils.JSONResponse(
		http.StatusOK,
		response,
		w,
	)
}
func getActivities(userid uint) []structs.Activity {
	db := database.SQL
	var activities []structs.Activity
	db.Where("user_id = ?", userid).Find(&activities)
	for index := 0; index < len(activities); index++ {
		activities[index].LoadIntervals()
	}
	spew.Dump(activities)
	return activities
}
