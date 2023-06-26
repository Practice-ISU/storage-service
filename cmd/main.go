package main

// "net/http"
import (
	"fmt"
	"log"
	"net"
	psql_cnf "storage-service/configs/postgre"
	folder_stor "storage-service/internal/adapters/db/folder"
	folder_server "storage-service/internal/adapters/grpc/folder"
	folder_service "storage-service/internal/domain/folder/service"
	folder_grpc "storage-service/pkg/grpc/folder"

	"google.golang.org/grpc"
)

func main() {
	psqlCnf := psql_cnf.GetConfig()
	storage := folder_stor.NewFileStorage(psqlCnf)
	service := folder_service.NewFolderService(storage)
	server := folder_server.NewFolderGrpcServer(service)

	s := grpc.NewServer()
	folder_grpc.RegisterFolderServiceServer(s, server)
	folderListener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("работаем")
	err = s.Serve(folderListener)
	if err != nil {
		log.Fatal(err)
	}
}
