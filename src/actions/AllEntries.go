package actions

import (
	"net/http"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// AllEntries get all user's entries
func AllEntries(w http.ResponseWriter, r *http.Request) {
	db := database.SQL
	var entries []structs.Entry
	usertoken := r.Context().Value("user").(structs.UserToken)
	db.Where("user_id = ?", usertoken.ID).Find(&entries)
	responseContent := structs.Token{
		Token: utils.GenerateToken(usertoken.ID, usertoken.Username, w),
		Data:  entries,
	}
	response := structs.Response{
		Message: "Ingesos cargados.",
		Content: responseContent,
	}
	utils.JSONResponse(http.StatusOK, response, w)
}
