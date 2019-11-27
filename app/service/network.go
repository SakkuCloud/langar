package service

import (
	"encoding/json"
	"errors"
	"langar/app/model"
	"langar/config"
	"langar/util"
	"net/http"
)

func GetNetworkList(filterKey string, filterValue string) (networks []model.Network, err error) {
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