package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	Host        string `yaml:"host" env-default:"localhost"`
	GrpcServers `yaml:"grpc_servers"`
	HttpServers `yaml:"http_servers"`
	Postgres    `yaml:"postgres"`
}

type GrpcServers struct {
	FileServer   `yaml:"file_server"`
	FolderServer `yaml:"folder_server"`
	PingServer   `yaml:"ping_server"`
}

type FileServer struct {
	Port string `yaml:"port" env-default:"8000"`
}

type FolderServer struct {
	Port string `yaml:"port" env-default:"8010"`
}

type PingServer struct {
	Port string `yaml:"port" env-default:"8020"`
}

type HttpServers struct {
	StorageServer `yaml:"storage_server"`
}

type StorageServer struct {
	Port        string        `yaml:"port" env-default:"8020"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeOut time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Postgres struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:"5432"`
	User     string `yaml:"user" env-default:"storage-user"`
	Database string `yaml:"database" env-default:"storage_database"`
	Password string `yaml:"password" env-default:"qwer"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
