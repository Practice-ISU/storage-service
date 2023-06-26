package folder

import (
	"context"
	folder_service "storage-service/internal/adapters/service/folder"
	"storage-service/internal/domain/folder/model"
	folder_grpc "storage-service/pkg/grpc/folder"
)

type folderGrpcServer struct {
	folder_grpc.UnimplementedFolderServiceServer
	service folder_service.FolderService
}

func NewFolderGrpcServer(service folder_service.FolderService) *folderGrpcServer {
	return &folderGrpcServer{
		service: service,
	}
}

func (server *folderGrpcServer) AddFolder(ctx context.Context, dto *folder_grpc.FolderAddDTO) (*folder_grpc.FolderResponse, error) {
	folderDto := &model.FolderAddDTO{
		UserId: dto.GetUserId(),
		FolderName: dto.GetFolderName(),
	}
	result, err := server.service.AddFolder(folderDto)
	if err != nil {
		return &folder_grpc.FolderResponse{
			Details: &folder_grpc.Details{
				Success: false,
				Mess: err.Error(),
			},
			Folder: nil,
		}, nil
	}
	return &folder_grpc.FolderResponse{
		Details: &folder_grpc.Details{
			Success: true,
			Mess: "",
		},
		Folder: &folder_grpc.FolderDTO{
			Id: result.Id,
			UserId: result.UserId,
			FolderName: result.FolderName,
		},
	}, nil
}
