package storage

import (
	"storage-service/pkg/clients/postgre"
)

type fileStorage struct {
	psqlClient *postgre.PsqlClient
}

func NewFileStorage(cnf postgre.Config) *fileStorage {
	client := postgre.NewPsqlCliennt(cnf)
	return &fileStorage{
		psqlClient: client,
	}
}

func (s *fileStorage) AddFile(filename string, folderId int64)     {}
func (s *fileStorage) DeleteFile(fileId int64)                     {}
func (s *fileStorage) RenameFile(fileId int64, newFileName string) {}

func (s *fileStorage) GetFile(filename string, folderId int64)           {}
func (s *fileStorage) GetFilesInStorage(filename string, folderId int64) {}
