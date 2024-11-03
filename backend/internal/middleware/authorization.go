package middlware

import (
	"net/http"
	"errors"

	"github.com/Molina-Q/svelte-wiki/backend/api"
	"github.com/Molina-Q/svelte-wiki/backend/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Authorization Middleware")
		// get the token from the request
		token := r.Header.Get("Authorization")
		log.Debug("Token: ", token)

		// check if the token is empty
		if token == "" {
			log.Error("Token is empty")
			api.ErrorResponse(w, UnAuthorizedError)
			return
		}

		// validate the token
		if !tools.ValidateToken(token) {
			log.Error("Token is invalid")
			api.ErrorResponse(w, UnAuthorizedError)
			return
		}

		// call the next handler
		next.ServeHTTP(w, r)
	})
}