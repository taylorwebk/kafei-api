package actions

import (
	"net/http"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// AllEntries get all user's entries
func AllEntries(w http.ResponseWriter, r *http.Request) {
	usertoken := r.Context().Value("user").(structs.UserToken)
	responseContent := structs.Token{
		Token: utils.GenerateToken(usertoken.ID, usertoken.Username, w),
		Data:  getEntries(usertoken.ID),
	}
	response := structs.Response{
		Message: "Ingesos cargados.",
		Content: responseContent,
	}
	utils.JSONResponse(http.StatusOK, response, w)
}
func getEntries(userid uint) []structs.Entry {
	db := database.SQL
	var entries []structs.Entry
	db.Where("user_id = ?", userid).Find(&entries)
	return entries
}
