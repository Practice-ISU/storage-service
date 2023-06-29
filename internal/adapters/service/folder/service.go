package folder

import "storage-service/internal/domain/folder/model"

type FolderService interface {
	AddFolder(dto *model.FolderAddDTO) (*model.FolderDTO, error)
	DeleteFolder(dto *model.FolderDeleteDTO) (*model.FolderDeleteResponse, error)
	GetFolder(dto *model.FolderGetByIdDTO) (*model.FolderDTO, error)
	GetAllFolders(dto *model.FolderGetAllDTO) (*[]model.FolderDTO, error)
	RenameFolder(dto *model.FolderRenameDTO) (*model.FolderDTO, error)
}
