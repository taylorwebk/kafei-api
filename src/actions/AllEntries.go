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
	entries, saving := GetEntries(usertoken.ID)
	response := structs.Response{
		Message: "Ingesos cargados.",
		Content: &struct {
			Entries []structs.EntryResponse `json:"entries"`
			Saving  float64                 `json:"saving"`
		}{
			Entries: entries,
			Saving:  saving,
		},
	}
	utils.JSONResponse(http.StatusOK, response, w)
}

// GetEntries get entries formated response
func GetEntries(userid uint) ([]structs.EntryResponse, float64) {
	db := database.SQL
	var entries []structs.Entry
	var entriesRes []structs.EntryResponse
	db.Where("user_id = ?", userid).Order("id desc").Find(&entries)
	sum := 0.0
	for _, entry := range entries {
		sum = sum + entry.Amount
		aux := structs.EntryResponse{}
		aux.ID = entry.ID
		aux.Date = utils.GetPrettyDate(entry.CreatedAt)
		aux.Time = utils.GetPrettyTime(entry.CreatedAt)
		aux.Literal = entry.Literal
		aux.Amount = entry.Amount
		entriesRes = append(entriesRes, aux)
	}
	return entriesRes, (sum / 10)
}
