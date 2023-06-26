package storage

import (
	"storage-service/internal/domain/file/model"
)

type FileStorage interface {
	AddFile(filename string, folderId int64) (*model.FileDTO, error)
	DeleteFile(fileId int64) (*model.FileDeleteResponceDTO, error)
	RenameFile(fileId int64, newFileName string) (*model.FileDTO, error)

	GetFileById(fileId int64) (*model.FileDTO, error)
	GetFileByFilename(filename string, folderId int64) (*model.FileDTO, error)
	GetFilesInStorage(folderId int64) (*[]model.FileDTO, error)
}
