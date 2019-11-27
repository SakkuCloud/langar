package handler

import (
	log "github.com/sirupsen/logrus"
	"langar/app/model"
	"net/http"
)

func IsAuthorized(w http.ResponseWriter, r *http.Request, auth *model.Auth) bool {
	if !auth.IsAuthorized(r.Header.Get("service"), r.Header.Get("service-key")) {
		log.Warnf("Unauthorized, %s", r.RemoteAddr)
		respondMessage(w, http.StatusUnauthorized, "Unauthorized")
		return false
	}
	return true
}