package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"langar/app/model"
	"langar/config"
	"langar/util"
	"net/http"
)

func GetNetworkList(filterKey string, filterValue string) (networks []model.NetworkDigest, err error) {
	client := &http.Client{Transport:util.GetUnixTransport(),}
	req, err := http.NewRequest(http.MethodGet, config.DockerNetworkAddress, nil)
	if err != nil {
		return
	}

	if filterKey != "" {
		if filterValue == "" {
			err = errors.New(filterKey+" value is empty")
			return
		}
		q := req.URL.Query()
		q.Add("filters", "{\""+filterKey+"\":{\""+filterValue+"\":true}}")
		req.URL.RawQuery = q.Encode()
	}

	rsp, err := client.Do(req)
	if err != nil || (rsp.StatusCode != http.StatusOK) {
		return
	}
	defer rsp.Body.Close()

	decoder := json.NewDecoder(rsp.Body)
	if err = decoder.Decode(&networks); err != nil {
		return
	}
	return
}

func GetNetworkById(id string) (network model.NetworkFull, err error) {
	client := &http.Client{Transport:util.GetUnixTransport(),}
	req, err := http.NewRequest(http.MethodGet, config.DockerNetworkAddress+"/"+id, nil)
	if err != nil {
		return
	}

	rsp, err := client.Do(req)
	if err != nil || (rsp.StatusCode != http.StatusOK) {
		return
	}
	defer rsp.Body.Close()

	decoder := json.NewDecoder(rsp.Body)
	if err = decoder.Decode(&network); err != nil {
		return
	}
	return
}

func CreateNetwork(network model.NetworkCreateReq) (networkRsp model.NetworkCreateResp, err error) {
	values, err := json.Marshal(network)
	if err != nil {
		return
	}

	client := &http.Client{Transport:util.GetUnixTransport(),}
	req, err := http.NewRequest(http.MethodPost, config.DockerNetworkAddress+"/create", bytes.NewBuffer(values))
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/json")
	rsp, err := client.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()

	decoder := json.NewDecoder(rsp.Body)
	if err = decoder.Decode(&networkRsp); err != nil {
		return
	}
	return
}