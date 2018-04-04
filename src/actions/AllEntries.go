package actions

import (
	"math"
	"net/http"
	"strconv"
	"time"

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
		aux.Date = getPrettyDate(entry.CreatedAt)
		aux.Time = getPrettyTime(entry.CreatedAt)
		aux.Literal = entry.Literal
		aux.Amount = entry.Amount
		entriesRes = append(entriesRes, aux)
	}
	return entriesRes, (sum / 10)
}
func getPrettyTime(date time.Time) string {
	now := time.Now()
	dayDiff := math.Abs(float64(now.Day() - date.Day()))
	if dayDiff <= 0 {
		minutesDiff := math.Abs(float64(now.Minute() - date.Minute()))
		hourDiff := math.Abs(float64(now.Hour() - date.Hour()))
		if hourDiff <= 1 {
			return "Hace " + strconv.Itoa(int(minutesDiff)) + " minuto(s)"
		}
		return "Hace " + strconv.Itoa(int(hourDiff)) + " hora(s)"
	}
	return date.Format("15:04")
}
func getPrettyDate(date time.Time) string {
	now := time.Now()
	dayDiff := math.Abs(float64(now.Day() - date.Day()))
	if dayDiff <= 5 {
		text := "Hoy"
		if dayDiff == 1 {
			text = "Ayer"
		}
		if dayDiff >= 2 {
			text = "Hace " + strconv.Itoa(int(dayDiff)) + " días"
		}
		return text
	}
	return date.Format("02-01-2006")
}
