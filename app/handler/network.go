package handler

import (
	"github.com/gorilla/mux"
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

	// call service
	networks, err := service.GetNetworkList(filterKey,filterValue)
	if err != nil {
		log.Warnf("Cannot fetch network list, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot fetch network list")
		return
	}

	log.Infof("Network list info sent")
	respondJSON(w, http.StatusOK, networks)
}

func GetNetwork(w http.ResponseWriter, r *http.Request) {
	networkId := mux.Vars(r)["id"]
	if networkId == "" {
		log.Warnf("Invalid request, empty container id")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty container id")
		return
	}

	network, err := service.GetNetworkById(networkId)
	if err != nil {
		log.Warnf("Cannot get network info, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot get network info")
		return
	}

	if network.Id == "" {
		log.Warnf("Network not found, %s", networkId)
		respondMessage(w, http.StatusNotFound, "Object not found")
		return
	}

	log.Infof("Network info sent, %s", network.Id)
	respondJSON(w, http.StatusOK, network)
}