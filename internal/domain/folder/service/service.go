package service

import (
	"fmt"
	"os"
	"storage-service/internal/domain/folder/model"
	"storage-service/internal/domain/folder/storage"
)

type config interface {
	GetStorageFolder() string
}


type folderService struct {
	storage storage.FolderStorage
	storageRoot string
}

func NewFolderService(storage storage.FolderStorage, cnf config) *folderService {
	return &folderService{
		storage: storage,
		storageRoot: cnf.GetStorageFolder(),
	}
}

func (s *folderService) getUserRoot(userId int64) string {
	return fmt.Sprintf("%s/user_%d/", s.storageRoot, userId)
}

func (s *folderService) ensureUserRootFolder(userId int64) (string, error) {
	folderName := s.getUserRoot(userId)
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err = os.Mkdir(folderName, 0666)
		if err != nil {
			return "", err
		}
	}

	return folderName, nil
}

func (s *folderService) AddFolder(dto *model.FolderAddDTO) (*model.FolderDTO, error) {
	root, err := s.ensureUserRootFolder(dto.UserId)
	if err != nil {
		return nil, err
	}
	result, _ := s.storage.GetFolderByName(&model.FolderGetByNameDTO{FolderName: dto.FolderName, UserId: dto.UserId})
	if result != nil {
		return nil, fmt.Errorf("this folder already exists")
	}

	err = os.Mkdir(root+dto.FolderName, 0666)
	if err != nil {
		return nil, err
	}
	return s.storage.AddFolder(dto)
}

func (s *folderService) DeleteFolder(dto *model.FolderDeleteDTO) (*model.FolderDeleteResponse, error) {
	result, err := s.storage.GetFolderById(&model.FolderGetByIdDTO{Id: dto.Id})
	if result == nil {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	root := s.getUserRoot(dto.UserId)
	err = os.RemoveAll(root + result.FolderName)
	if err != nil {
		return nil, err
	}
	return s.storage.DeleteFolder(dto)
}

func (s *folderService) RenameFolder(dto *model.FolderRenameDTO) (*model.FolderDTO, error) {
	result, err := s.storage.GetFolderById(&model.FolderGetByIdDTO{Id: dto.Id})
	if result == nil {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	root := s.getUserRoot(dto.UserId)
	err = os.Rename(root+result.FolderName, root+dto.NewName)
	if err != nil {
		return nil, err
	}
	return s.storage.RenameFolder(dto)
}

func (s *folderService) GetUrl(dto *model.FolderGetByIdDTO) (string, error) {
	folder, err := s.storage.GetFolderById(dto)
	if err != nil {
		return "", nil
	}
	url := s.getUserRoot(folder.UserId)
	return url + folder.FolderName + "/", nil
}

func (s *folderService) GetFolder(dto *model.FolderGetByIdDTO) (*model.FolderDTO, error) {
	return s.storage.GetFolderById(dto)
}

func (s *folderService) GetAllFolders(dto *model.FolderGetAllDTO) (*[]model.FolderDTO, error) {
	return s.storage.GetAllFolders(dto)
}
