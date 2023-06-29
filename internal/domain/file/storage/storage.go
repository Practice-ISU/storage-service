package storage

import (
	"storage-service/internal/domain/file/model"
)

type FileStorage interface {
	AddFile(dto *model.FileAddDTO) (*model.FileDTO, error)
	DeleteFile(dto *model.FileDeleteDTO) (*model.FileDeleteResponceDTO, error)
	RenameFile(dto *model.FileRenameDTO) (*model.FileDTO, error)

	GetFileById(dto *model.FileGetByIdDTO) (*model.FileDTO, error)
	GetFileByFilename(dto *model.FileGetByNameDTO) (*model.FileDTO, error)
	GetFilesInFolder(dto *model.FileGetAllDTO) (*[]model.FileDTO, error)
}
