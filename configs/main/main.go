package mainconfig

import "os"

type mainConfig struct {
	ipAddr            string
	filePort          string
	folderPort        string
	folderServiceName string
	fileServiceName   string
	storageFolder     string
	storagePort       string
}

func GetMainConfig() *mainConfig {
	conf := &mainConfig{
		ipAddr:            "127.0.0.1",
		filePort:          "8080",
		folderPort:        "8000",
		folderServiceName: "FolderGrpcService",
		fileServiceName:   "FileGrpcService",
		storageFolder:     "./storage",
		storagePort:       "9000",
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
	folderServiceName := os.Getenv("FOLDER_SERVICE_NAME")
	if ip != "" {
		conf.folderServiceName = folderServiceName
	}
	fileServiceName := os.Getenv("FILE_SERVICE_NAME")
	if ip != "" {
		conf.fileServiceName = fileServiceName
	}
	storageFolder := os.Getenv("STORAGE_FOLDER")
	if ip != "" {
		conf.storageFolder = storageFolder
	}
	storagePort := os.Getenv("STORAGE_PORT")
	if storageFolder != "" {
		conf.storagePort = storagePort
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

func (cnf *mainConfig) GetFolderServiceName() string {
	return cnf.folderServiceName
}

func (cnf *mainConfig) GetFileServiceName() string {
	return cnf.fileServiceName
}

func (cnf *mainConfig) GetStorageFolder() string {
	return cnf.storageFolder
}

func (cnf *mainConfig) GetStoragePort() string {
	return cnf.storagePort
}
