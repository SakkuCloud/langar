package config

import "time"

const (
	SentryTimeout time.Duration = 10

	DockerAPIVersion     string = "1.41"
	DockerNetworkAddress string = "http://" + DockerAPIVersion + "/networks"

	InvalidArgsMessage string = "invalid args"
)

var NetworkFilterMap = map[string]string{"d": "driver"}
var StartTime time.Time

var Config = struct {
	Port    string `default:"3000"`
	LogFile string `default:"/var/log/langar.log"`

	SentryDSN string

	AccessKey string `required:"true"`
	SecretKey string `required:"true"`

	Docker struct {
		Socket string `default:"/run/docker.sock"`
	}
}{}
