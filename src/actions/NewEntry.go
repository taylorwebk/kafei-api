package actions

import (
	"encoding/json"
	"net/http"

	"github.com/taylorwebk/kafei-api/src/database"
	"github.com/taylorwebk/kafei-api/src/structs"
	"github.com/taylorwebk/kafei-api/src/utils"
)

// NewEntry register a new entry to database
func NewEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	db := database.SQL
	user := structs.User{}
	entry := structs.Entry{}
	err = json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		utils.JSONResponse(
			http.StatusUnprocessableEntity,
			structs.Response{
				Message: "Error en los datos recibidos",
			},
			w,
		)
		return
	}
	usertoken := r.Context().Value("user").(structs.UserToken)
	db.Where("username = ?", usertoken.Username).First(&user)
	user.AddEntry(entry)
	db.Save(&user)
	entries, saving := GetEntries(usertoken.ID)
	response := structs.Response{
		Message: "Nuevo Ingreso Guardado.",
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
