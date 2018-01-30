package middlewares

import (
	"net/http"

	"github.com/taylorwebk/kafei-api/src/structs"

	"context"

	"github.com/taylorwebk/kafei-api/src/utils"
)

// Login manages the login
func Login(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenstring := r.Header.Get("Authorization")
	username, id := utils.GetUsernameByToken(tokenstring, w)
	if username != "error" {
		ctx := context.WithValue(r.Context(), "user", structs.UserToken{
			Username: username,
			ID:       id,
		})
		//ctx2 := context.WithValue(r.Context(), "id", id)
		next(w, r.WithContext(ctx))
	} else {
		return
	}
}
