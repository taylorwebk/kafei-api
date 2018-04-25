package actions

import (
	"net/http"

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
func getActivities(userid uint) []structs.ActivityResponse {
	db := database.SQL
	var activities []structs.Activity
	var activityRes []structs.ActivityResponse
	db.Where("user_id = ?", userid).Find(&activities)
	for _, activity := range activities {
		aux := structs.ActivityResponse{}
		aux.ID = activity.ID
		aux.Literal = activity.Literal
		aux.Time = utils.GetPrettyTime(activity.CreatedAt)
		aux.Date = utils.GetPrettyDate(activity.CreatedAt)
		aux.Hours, aux.Mins = getTotal(activity.Intervs)
		activityRes = append(activityRes, aux)
	}
	return activityRes
}
func getTotal(intervs []structs.Interv) (float64, float64) {
	totalH := 0.0
	totalM := 0.0
	for _, interv := range intervs {
		totalH += (interv.End.Sub(interv.End)).Hours()
		totalH += (interv.End.Sub(interv.End)).Minutes()
	}
	return totalH, totalM
}
