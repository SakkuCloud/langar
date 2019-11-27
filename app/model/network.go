package model

type Network struct {
	Id         string                          `json:"Id"`
	Name       string                          `json:"Name"`
	Created    string                          `json:"Created"`
	Scope      string                          `json:"Scope"`
	Driver     string                          `json:"Driver"`
	IPAM       NetworkIPAM                     `json:"IPAM"`
	Internal   bool                            `json:"Internal"`
	Attachable bool                            `json:"Attachable"`
	Ingress    bool                            `json:"Ingress"`
	Containers map[string]NetworkContainerInfo `json:"Containers"`
	Options    NetworkOptions                  `json:"Options"`
}

type NetworkDigest struct {
	Id     string      `json:"Id"`
	Name   string      `json:"Name"`
	Scope  string      `json:"Scope"`
	Driver string      `json:"Driver"`
	IPAM   NetworkIPAM `json:"IPAM"`
}

type NetworkIPAM struct {
	Driver string              `json:"Driver"`
	Config []NetworkIPAMConfig `json:"Config"`
}

type NetworkIPAMConfig struct {
	Subnet  string `json:"Subnet"`
	Gateway string `json:"Gateway"`
}

type NetworkContainerInfo struct {
	Name        string `json:"Name"`
	EndpointId  string `json:"EndpointID"`
	MacAddress  string `json:"MacAddress"`
	IPv4Address string `json:"IPv4Address"`
}

type NetworkOptions struct {
	BridgeDefaultBridge   string `json:"com.docker.network.bridge.default_bridge"`
	BridgeIPMasquerade    string `json:"com.docker.network.bridge.enable_ip_masquerade"`
	BridgeHostBindingIPv4 string `json:"com.docker.network.bridge.host_binding_ipv4"`
	BridgeName            string `json:"com.docker.network.bridge.name"`
	DriverMTU             string `json:"com.docker.network.driver.mtu"`
}
