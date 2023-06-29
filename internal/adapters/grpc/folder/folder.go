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
		UserId:     dto.GetUserId(),
		FolderName: dto.GetFolderName(),
	}
	result, err := server.service.AddFolder(folderDto)
	if err != nil {
		return &folder_grpc.FolderResponse{
			Details: &folder_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			Folder: nil,
		}, nil
	}
	return &folder_grpc.FolderResponse{
		Details: &folder_grpc.Details{
			Success: true,
			Mess:    "",
		},
		Folder: &folder_grpc.FolderDTO{
			Id:         result.Id,
			UserId:     result.UserId,
			FolderName: result.FolderName,
		},
	}, nil
}

func (server *folderGrpcServer) DeleteFolder(ctx context.Context, dto *folder_grpc.FolderDeleteDTO) (*folder_grpc.Details, error) {
	folderDto := &model.FolderDeleteDTO{
		Id:     dto.Id,
		UserId: dto.UserId,
	}
	result, err := server.service.DeleteFolder(folderDto)
	if err != nil {
		return &folder_grpc.Details{
			Success: false,
			Mess:    err.Error(),
		}, nil
	}
	return &folder_grpc.Details{
		Success: result.Success,
		Mess:    result.Mess,
	}, nil
}

func (server *folderGrpcServer) GetFolder(ctx context.Context, dto *folder_grpc.FolderGetDTO) (*folder_grpc.FolderResponse, error) {
	folderDto := &model.FolderGetByIdDTO{
		Id:     dto.Id,
	}

	result, err := server.service.GetFolder(folderDto)
	if err != nil {
		return &folder_grpc.FolderResponse{
			Details: &folder_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			Folder: nil,
		}, nil
	}
	return &folder_grpc.FolderResponse{
		Details: &folder_grpc.Details{
			Success: true,
			Mess:    "",
		},
		Folder: &folder_grpc.FolderDTO{
			Id:         result.Id,
			UserId:     result.UserId,
			FolderName: result.FolderName,
		},
	}, nil
}

func (server *folderGrpcServer) GetAllFolders(ctx context.Context, dto *folder_grpc.UserId) (*folder_grpc.AllFoldersResponse, error) {
	foldersDto := &model.FolderGetAllDTO{
		UserId: dto.UserId,
	}

	result, err := server.service.GetAllFolders(foldersDto)
	folders := make([]*folder_grpc.FolderDTO, len(*result))
	for i, folder := range *result {
		folders[i] = &folder_grpc.FolderDTO{
			Id:         folder.Id,
			FolderName: folder.FolderName,
			UserId:     folder.UserId,
		}
	}
	if err != nil {
		return &folder_grpc.AllFoldersResponse{
			Details: &folder_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			Folders: nil,
		}, nil
	}
	return &folder_grpc.AllFoldersResponse{
		Details: &folder_grpc.Details{
			Success: true,
			Mess:    "",
		},
		Folders: folders,
	}, nil
}

func (server *folderGrpcServer) RenameFolder(ctx context.Context, dto *folder_grpc.FolderRenameDTO) (*folder_grpc.FolderResponse, error) {
	folderDto := &model.FolderRenameDTO{
		Id:      dto.Id,
		NewName: dto.NewFolderName,
		UserId:  dto.UserId,
	}

	result, err := server.service.RenameFolder(folderDto)
	if err != nil {
		return &folder_grpc.FolderResponse{
			Details: &folder_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			Folder: nil,
		}, nil
	}
	return &folder_grpc.FolderResponse{
		Details: &folder_grpc.Details{
			Success: true,
			Mess:    "",
		},
		Folder: &folder_grpc.FolderDTO{
			Id:         result.Id,
			UserId:     result.UserId,
			FolderName: result.FolderName,
		},
	}, nil
}
