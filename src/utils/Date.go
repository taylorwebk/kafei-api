package utils

import (
	"math"
	"strconv"
	"time"
)

// GetPrettyTime pretty time
func GetPrettyTime(date time.Time) string {
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

// GetPrettyDate pretty date
func GetPrettyDate(date time.Time) string {
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
