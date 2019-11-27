package model

type Network struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	Created string `json:"Created"`
	Scope string `json:"Scope"`
	Driver string `json:"Driver"`
	Internal bool `json:"Internal"`
	Attachable bool `json:"Attachable"`
	Containers Container `json:"Containers"`
}
