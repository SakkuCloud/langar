package util

import (
	"context"
	"langar/config"
	"net"
	"net/http"
)

func GetUnixTransport() *http.Transport {
	return &http.Transport{
		DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("unix", config.Config.Docker.Host)
		},
	}
}
