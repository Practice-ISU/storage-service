package mainconfig

import "os"

type mainConfig struct {
	ipAddr      string
	filePort    string
	folderPort  string
	serviceName string
}

func GetMainConfig() *mainConfig {
	conf := &mainConfig{
		ipAddr:      "127.0.0.1",
		filePort:    "8080",
		folderPort:  "8000",
		serviceName: "StorageGrpcServer",
	}
	ip := os.Getenv("SERVICE_IP")
	if ip != "" {
		conf.ipAddr = ip
	}
	filePort := os.Getenv("FILE_PORT")
	if ip != "" {
		conf.filePort = filePort
	}
	folderPort := os.Getenv("FOLDER_PORT")
	if ip != "" {
		conf.folderPort = folderPort
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if ip != "" {
		conf.serviceName = serviceName
	}
	return conf
}

func (cnf *mainConfig) GetIp() string {
	return cnf.ipAddr
}

func (cnf *mainConfig) GetFilePort() string {
	return cnf.filePort
}

func (cnf *mainConfig) GetFolderPort() string {
	return cnf.folderPort
}

func (cnf *mainConfig) GetServiceName() string {
	return cnf.serviceName
}