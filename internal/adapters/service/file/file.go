package file

import "storage-service/internal/domain/file/model"

type FileService interface {
	AddFile(dto *model.FileAddDTO) (*model.FileDTO, error)
	DeleteFile(dto *model.FileDeleteDTO) (*model.FileDeleteResponceDTO, error)
	RenameFile(dto *model.FileRenameDTO) (*model.FileDTO, error)

	GetFileById(dto *model.FileGetByIdDTO) (*model.FileDTO, error)
	GetFilesInFolder(dto *model.FileGetAllDTO) (*[]model.FileDTO, error)

	GetFileBase64(dto *model.FileGetByIdDTO) (*model.FileBase64, error)
	GetFilesInFolderBase64(dto *model.FileGetAllDTO) (*[]model.FileBase64, error)
	GetFilesInFolderZip(dto *model.FileGetAllDTO) (*model.FilesZip, error)
}
