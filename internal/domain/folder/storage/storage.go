package storage

import (
	"storage-service/internal/domain/folder/model"
)

type FolderStorage interface {
	AddFolder(dto *model.FolderAddDTO) (*model.FolderDTO, error)
	DeleteFolder(dto *model.FolderDeleteDTO) (*model.FolderDeleteResponse, error)
	GetAllFolders(dto *model.FolderGetAllDTO) (*[]model.FolderDTO, error)
	GetFolderById(dto *model.FolderGetByIdDTO) (*model.FolderDTO, error)
	GetFolderByName(dto *model.FolderGetByNameDTO) (*model.FolderDTO, error)
	RenameFolder(dto *model.FolderRenameDTO) (*model.FolderDTO, error)


}
