package main

import (
	"fmt"
	"log"
	"net"

	mainconfig "storage-service/configs/main"
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
)

func main() {
	psqlCnf := psql_cnf.GetConfig()
	mainCnf := mainconfig.GetMainConfig()

	folderStorage := folder_stor.NewFileStorage(psqlCnf)
	fileStorage := file_stor.NewFileStorage(psqlCnf)

	folderService := folder_service.NewFolderService(folderStorage)
	fileService := file_service.NewFileService(fileStorage, folderService)
	folderServer := folder_server.NewFolderGrpcServer(folderService)
	fileServer := file_server.NewFileGrpcServer(fileService)

	s := grpc.NewServer()
	folder_grpc.RegisterFolderServiceServer(s, folderServer)
	file_grpc.RegisterFileServiceServer(s, fileServer)

	folderListener, err := net.Listen("tcp", ":" + mainCnf.GetFolderPort())
	if err != nil {
		log.Fatal(err)
	}
	fileListener, err := net.Listen("tcp", ":" + mainCnf.GetFilePort())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = s.Serve(fileListener)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("File service is working on: " + mainCnf.GetIp() + ":" + mainCnf.GetFilePort())
	}()

	fmt.Println("Folder service is working on: " + mainCnf.GetIp() + ":" + mainCnf.GetFolderPort())
	err = s.Serve(folderListener)
	if err != nil {
		log.Fatal(err)
	}
}
