package handler

import (
	log "github.com/sirupsen/logrus"
	"langar/app/service"
	"langar/config"
	"net/http"
)

func GetNetworks(w http.ResponseWriter, r *http.Request) {
	// check query params
	filterKey, filterValue := "",""
	queries := r.URL.Query()
	if len(queries) > 0 {
		if len(queries) > 1 {
			log.Warnf("Invalid query params")
			respondMessage(w, http.StatusBadRequest, "Invalid query params")
			return
		}

		queryMatched := false
		for k, v := range queries{
			if value, ok := config.NetworkFilterMap[k]; ok {
				filterKey, filterValue = value, v[0]
				queryMatched = true
			}
		}
		if !queryMatched {
			log.Warnf("Invalid query params")
			respondMessage(w, http.StatusBadRequest, "Invalid query params")
			return
		}
	}

	networks, err := service.GetNetworkList(filterKey,filterValue)
	if err != nil {
		log.Warnf("Cannot fetch network list, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot fetch network list")
		return
	}

	respondJSON(w, http.StatusOK, networks)
}
