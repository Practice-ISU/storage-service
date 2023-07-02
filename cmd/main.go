package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	discovery_config "storage-service/configs/grpc/discovery"
	main_config "storage-service/configs/main"
	psql_cnf "storage-service/configs/postgre"

	"google.golang.org/grpc"

	file_stor "storage-service/internal/adapters/db/file"
	folder_stor "storage-service/internal/adapters/db/folder"
	file_server "storage-service/internal/adapters/grpc/file"
	folder_server "storage-service/internal/adapters/grpc/folder"
	file_service "storage-service/internal/domain/file/service"
	folder_service "storage-service/internal/domain/folder/service"
	file_grpc "storage-service/pkg/grpc/file"
	folder_grpc "storage-service/pkg/grpc/folder"

	ping_server "storage-service/pkg/api/grpc/ping"
	ping_grpc "storage-service/pkg/grpc/discovery/ping"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func fileServerWithCors(dir string) http.Handler {
	fs := http.FileServer(http.Dir(dir))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		fs.ServeHTTP(w, r)
	})
}

func main() {
	psqlCnf := psql_cnf.GetConfig()
	mainCnf := main_config.GetMainConfig()
	discoveryCnf := discovery_config.GetConfig()
	serverOptions := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(1024 * 1024 * 1024),
		grpc.MaxSendMsgSize(1024 * 1024 * 1024),
	}

	folderStorage := folder_stor.NewFileStorage(psqlCnf)
	fileStorage := file_stor.NewFileStorage(psqlCnf)

	folderService := folder_service.NewFolderService(folderStorage, mainCnf)
	fileService := file_service.NewFileService(fileStorage, folderService, mainCnf)
	folderServer := folder_server.NewFolderGrpcServer(folderService)
	fileServer := file_server.NewFileGrpcServer(fileService, mainCnf)



	pingServer, err := ping_server.NewDiscoveryPingServer(discoveryCnf, mainCnf)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(serverOptions...)
	folder_grpc.RegisterFolderServiceServer(s, folderServer)
	file_grpc.RegisterFileServiceServer(s, fileServer)
	ping_grpc.RegisterDiscoveryPingServer(s, pingServer)

	folderListener, err := net.Listen("tcp", ":"+mainCnf.GetFolderPort())
	if err != nil {
		log.Fatal(err)
	}
	fileListener, err := net.Listen("tcp", ":"+mainCnf.GetFilePort())
	if err != nil {
		log.Fatal(err)
	}

	pingListener, err := net.Listen("tcp", ":"+discoveryCnf.GetPingPort())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		fmt.Println("Ping service is working on: " + mainCnf.GetIp() + ":" + discoveryCnf.GetPingPort())
		err = s.Serve(pingListener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		fmt.Println("File service is working on: " + mainCnf.GetIp() + ":" + mainCnf.GetFilePort())
		err = s.Serve(fileListener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		pingServer.SendRegistrationRequest()
	}()

	go func() {
		for {
			pingServer.StartTimeout(pingServer.SendRegistrationRequest)
			time.Sleep(6 * time.Minute)
		}
	}()

	go func() {
		// fs := http.FileServer(http.Dir())
		// http.Handle("/storage/", http.StripPrefix("/storage/", fs))

		// c := cors.New(cors.Options{
		// 	AllowedOrigins: []string{"*"},
		// 	AllowedMethods: []string{"*"},
		// 	AllowedHeaders: []string{"*"},
		// })

		// handler := c.Handler(http.DefaultServeMux)

		log.Println("Static server is working on: localhost:" + mainCnf.GetStoragePort() + mainCnf.GetStorageFolder())

		http.Handle("/storage/", http.StripPrefix("/storage/", fileServerWithCors(mainCnf.GetStorageFolder())))
		err = http.ListenAndServe(":"+mainCnf.GetStoragePort(), nil)
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Folder service is working on: " + mainCnf.GetIp() + ":" + mainCnf.GetFolderPort())
	err = s.Serve(folderListener)
	if err != nil {
		log.Fatal(err)
	}
}
