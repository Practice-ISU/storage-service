package file

import (
	"context"
	file_service "storage-service/internal/adapters/service/file"
	"storage-service/internal/domain/file/model"
	file_grpc "storage-service/pkg/grpc/file"
)

type config interface {
	GetIp() string
}

type fileGrpcServer struct {
	file_grpc.UnimplementedFileServiceServer
	service file_service.FileService
	ip string
}

func NewFileGrpcServer(service file_service.FileService, cnf config) *fileGrpcServer {
	return &fileGrpcServer{
		service: service,
		ip: cnf.GetIp(),
	}
}

func (server *fileGrpcServer) AddFile(ctx context.Context, dto *file_grpc.FileAddDTO) (*file_grpc.FileResponse, error) {
	result, err := server.service.AddFile(&model.FileAddDTO{
		FolderId: dto.FolderId,
		FileName: dto.FileName,
		Base64:   dto.Base64,
		Format:   dto.Format,
	})

	if err != nil {
		return &file_grpc.FileResponse{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			File: nil,
		}, nil
	}

	return &file_grpc.FileResponse{
		Details: &file_grpc.Details{
			Success: true,
			Mess:    "",
		},
		File: &file_grpc.FileDTO{
			Id:       result.Id,
			FolderId: result.FolderId,
			Filename: result.FileName,
			Url: "http://" + server.ip + result.Url,
		},
	}, nil
}

func (server *fileGrpcServer) DeleteFile(ctx context.Context, dto *file_grpc.FileDeleteDTO) (*file_grpc.Details, error) {
	result, err := server.service.DeleteFile(&model.FileDeleteDTO{Id: dto.Id})
	if err != nil {
		return &file_grpc.Details{
			Success: false,
			Mess:    err.Error(),
		}, nil
	}

	return &file_grpc.Details{
		Success: result.Success,
		Mess:    result.Mess,
	}, nil
}

func (server *fileGrpcServer) RenameFile(ctx context.Context, dto *file_grpc.FileRenameDTO) (*file_grpc.FileResponse, error) {
	result, err := server.service.RenameFile(&model.FileRenameDTO{Id: dto.GetId(), NewFileName: dto.GetNewFilename()})

	if err != nil {
		return &file_grpc.FileResponse{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			File: nil,
		}, nil
	}

	return &file_grpc.FileResponse{
		Details: &file_grpc.Details{
			Success: true,
			Mess:    "",
		},
		File: &file_grpc.FileDTO{
			Id:       result.Id,
			FolderId: result.FolderId,
			Filename: result.FileName,
			Url: "http://" + server.ip + result.Url,
		},
	}, nil
}

func (server *fileGrpcServer) GetAllFilesInFolder(ctx context.Context, dto *file_grpc.FileGetAllDTO) (*file_grpc.FileAllResponse, error) {
	result, err := server.service.GetFilesInFolder(&model.FileGetAllDTO{FolderId: dto.FolderId})
	if err != nil {
		return &file_grpc.FileAllResponse{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			Files: nil,
		}, nil
	}
	files := make([]*file_grpc.FileDTO, len(*result))
	for i, file := range *result {
		files[i] = &file_grpc.FileDTO{
			Id:       file.Id,
			FolderId: file.FolderId,
			Filename: file.FileName,
			Url: "http://" + server.ip + file.Url,
		}
	}
	return &file_grpc.FileAllResponse{
		Details: &file_grpc.Details{
			Success: true,
			Mess:    "",
		},
		Files: files,
	}, nil
}

func (server *fileGrpcServer) GetFile(ctx context.Context, dto *file_grpc.FileGetDTO) (*file_grpc.FileResponse, error) {
	result, err := server.service.GetFileById(&model.FileGetByIdDTO{Id: dto.Id})
	if err != nil {
		return &file_grpc.FileResponse{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			File: nil,
		}, nil
	}

	return &file_grpc.FileResponse{
		Details: &file_grpc.Details{
			Success: true,
			Mess:    "",
		},
		File: &file_grpc.FileDTO{
			Id:       result.Id,
			FolderId: result.FolderId,
			Filename: result.FileName,
			Url: "http://" + server.ip + result.Url,
		},
	}, nil
}

func (server *fileGrpcServer) GetFileBase64(ctx context.Context, dto *file_grpc.FileGetDTO) (*file_grpc.FileBase64Response, error) {
	result, err := server.service.GetFileBase64(&model.FileGetByIdDTO{Id: dto.Id})
	if err != nil {
		return &file_grpc.FileBase64Response{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			File: nil,
		}, nil
	}

	return &file_grpc.FileBase64Response{
		Details: &file_grpc.Details{
			Success: true,
			Mess:    "",
		},
		File: &file_grpc.FileBase64{
			Id:       result.Id,
			FolderId: result.FolderId,
			FileName: result.FileName,
			Base64:   result.Base64,
		},
	}, nil
}

func (server *fileGrpcServer) GetAllFilesInFolderBase64(ctx context.Context, dto *file_grpc.FileGetAllDTO) (*file_grpc.FileAllBase64Response, error) {
	result, err := server.service.GetFilesInFolderBase64(&model.FileGetAllDTO{FolderId: dto.FolderId})
	if err != nil {
		return &file_grpc.FileAllBase64Response{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			Files: nil,
		}, nil
	}

	files := make([]*file_grpc.FileBase64, len(*result))
	for i, file := range *result {
		files[i] = &file_grpc.FileBase64{
			Id:       file.Id,
			FolderId: file.FolderId,
			FileName: file.FileName,
			Base64:   file.Base64,
		}
	}
	return &file_grpc.FileAllBase64Response{
		Details: &file_grpc.Details{
			Success: true,
			Mess:    "",
		},
		Files: files,
	}, nil
}

func (server *fileGrpcServer) GetAllFilesInFolderZip(ctx context.Context, dto *file_grpc.FileGetAllDTO) (*file_grpc.FileAllZipResponse, error) {
	result, err := server.service.GetFilesInFolderZip(&model.FileGetAllDTO{FolderId: dto.FolderId})

	if err != nil {
		return &file_grpc.FileAllZipResponse{
			Details: &file_grpc.Details{
				Success: false,
				Mess:    err.Error(),
			},
			ZipName: "",
			Zip:     []byte{},
		}, nil
	}

	return &file_grpc.FileAllZipResponse{
		Details: &file_grpc.Details{
			Success: false,
			Mess:    "",
		},
		ZipName: result.ZipName,
		Zip:     result.Data,
	}, nil

}
