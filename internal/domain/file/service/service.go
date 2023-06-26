package service

import (
	"fmt"
	"storage-service/internal/domain/file/model"
	"storage-service/internal/domain/file/storage"
)

type fileService struct {
	storage storage.FileStorage
}

func NewFileService(storage storage.FileStorage) *fileService {
	return &fileService{
		storage: storage,
	}
}


func (s *fileService) AddFile(filename string, folderId int64) (*model.FileDTO, error) {
	result, _ := s.GetFileByFilename(filename, folderId)
	if result != nil {
		return nil, fmt.Errorf("this file in this folder is already exists")
	}
	return s.storage.AddFile(filename, folderId)

}
func (s *fileService) DeleteFile(fileId int64) (*model.FileDeleteResponceDTO, error) {
	result, _ := s.GetFileById(fileId)
	if result == nil {
		return nil, fmt.Errorf("this file is not exist")
	}
	return nil, nil
}
func (s *fileService) RenameFile(fileId int64, newFileName string) (*model.FileDTO, error) {
	return nil, nil
}

func (s *fileService) GetFileById(fileId int64) (*model.FileDTO, error) {
	return nil, nil
}
func (s *fileService) GetFileByFilename(filename string, folderId int64) (*model.FileDTO, error) {
	return nil, nil
}
func (s *fileService) GetFilesInStorage(filename string, folderId int64) ([]*model.FileDTO, error) {
	return nil, nil
}

