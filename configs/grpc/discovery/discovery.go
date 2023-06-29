package discovery

import "os"

type config struct {
	pingPort string
	discoveryAddr string
}

func GetConfig() *config {
	conf := &config{
		pingPort: "8010",
		discoveryAddr: "158.160.26.1:80",
	}
	pingPort := os.Getenv("PORT_PING")
	if pingPort != "" {
		conf.pingPort = pingPort
	}

	discoveryAddr := os.Getenv("DISCOVERY_ADDR")
	if discoveryAddr != "" {
		conf.discoveryAddr = discoveryAddr
	}
	return conf
}

func (cnf *config) GetPingPort() string {
	return cnf.pingPort
}

func (cnf *config) GetDiscoveryAddr() string {
	return cnf.discoveryAddr
}