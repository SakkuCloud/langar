package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"langar/app/model"
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
		log.Warnf("Invalid request, empty network id")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty network id")
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

func CreateNetwork(w http.ResponseWriter, r *http.Request) {
	network := model.NetworkCreateReq{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&network); err != nil {
		log.Warnf("Cannot decode network object in CreateNetwork, %s", err.Error())
		respondMessage(w, http.StatusBadRequest, "Cannot decode object")
		return
	}
	defer r.Body.Close()

	rsp, err := service.CreateNetwork(network)
	if err != nil {
		log.Warnf("Cannot create network, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot create network")
		return
	}

	if rsp.Id == "" {
		log.Warnf("Cannot create network, %s", rsp.Message)
		respondMessage(w, http.StatusBadRequest, rsp.Message)
		return
	}

	log.Infof("Network created, %s", rsp.Id)
	respondJSON(w, http.StatusCreated, rsp)
}

func DeleteNetwork(w http.ResponseWriter, r *http.Request) {
	networkId := mux.Vars(r)["id"]
	if networkId == "" {
		log.Warnf("Invalid request, empty network id")
		respondMessage(w, http.StatusBadRequest, "Invalid request, empty network id")
		return
	}

	rsp, err := service.DeleteNetworkById(networkId)
	if err != nil {
		log.Warnf("Cannot delete network, %s", err.Error())
		respondMessage(w, http.StatusInternalServerError, "Cannot delete network")
		return
	}

	if rsp == http.StatusNotFound {
		log.Warnf("Cannot delete network, no such network")
		respondMessage(w, http.StatusNotFound , "Cannot delete network, no such network")
		return
	}

	if rsp == http.StatusForbidden {
		log.Warnf("Cannot delete network, operation not supported for pre-defined networks")
		respondMessage(w, http.StatusBadRequest, "Cannot delete network, operation not supported for pre-defined networks")
		return
	}

	if rsp == http.StatusInternalServerError {
		log.Warnf("something went wrong")
		respondMessage(w, http.StatusInternalServerError, "Cannot delete network, something went wrong")
		return
	}

	log.Infof("Network deleted, %s", networkId)
	respondJSON(w, http.StatusNoContent, nil)
}