package storage

import (
	"storage-service/internal/domain/folder/model"
)

type FolderStorage interface {
	AddFolder(dto *model.FolderAddDTO) (*model.FolderDTO, error)
	DeleteFolder(dto *model.FolderDeleteDTO) (*model.FolderDeleteResponse, error)
	GetAllFolders(userId int64) (*[]model.FolderDTO, error)
	GetFolderById(dto *model.FolderGetDTO) (*model.FolderDTO, error)
	GetFolderByName(dto *model.FolderGetDTO) (*model.FolderDTO, error)
	RenameFolder(dto *model.FolderRenameDTO) (*model.FolderDTO, error)
}
