package storage

import (
	"fmt"
	"storage-service/internal/domain/file/model"
	"storage-service/pkg/clients/postgre"
)

type scanner interface {
	Scan(dest ...any) error
}

func scanFile(scanner scanner) (*model.FileDTO, error) {
	file := &model.FileDTO{}
	err := scanner.Scan(&file.Id, &file.FolderId, &file.FileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

type fileStorage struct {
	psqlClient *postgre.PsqlClient
}

func NewFileStorage(cnf postgre.Config) *fileStorage {
	client := postgre.NewPsqlCliennt(cnf)
	return &fileStorage{
		psqlClient: client,
	}
}

func (s *fileStorage) getFileWithWhereCase(whereCase string) (*model.FileDTO, error) {
	stmt := "SELECT id, folder_id, filename FROM files"
	if whereCase != "" {
		stmt += " WHERE " + whereCase
	}
	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result := db.QueryRow(stmt)
	return scanFile(result)
}

func (s *fileStorage) getFilesWithWhereCase(whereCase string) (*[]model.FileDTO, error) {
	stmt := "SELECT id, folder_id, filename FROM files"
	if whereCase != "" {
		stmt += " WHERE " + whereCase
	}
	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	files := []model.FileDTO{}
	for result.Next() {
		file, err := scanFile(result)
		if err != nil {
			return nil, err
		}
		files = append(files, *file)
	}
	return &files, nil
}

func (s *fileStorage) AddFile(dto *model.FileAddDTO) (*model.FileDTO, error) {
	stmt := "INSERT INTO files (folder_id, filename) VALUES" + dto.ExtractInsertSQL()
	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var id int64
	err = db.QueryRow(stmt + " RETURNING id;").Scan(&id)

	if err != nil {
		return nil, err
	}

	return &model.FileDTO{
		Id:       id,
		FolderId: dto.FolderId,
		FileName: dto.FileName,
	}, nil
}
func (s *fileStorage) DeleteFile(dto *model.FileDeleteDTO) (*model.FileDeleteResponceDTO, error) {
	stmt := fmt.Sprintf("DELETE FROM files WHERE id = %d", dto.Id)
	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec(stmt)

	if err != nil {
		return &model.FileDeleteResponceDTO{
			Success: false,
			Mess:    err.Error(),
		}, err
	}
	return &model.FileDeleteResponceDTO{
		Success: true,
		Mess:    "",
	}, nil
}
func (s *fileStorage) RenameFile(dto *model.FileRenameDTO) (*model.FileDTO, error) {
	stmt := fmt.Sprintf("UPDATE files SET filename = '%s' WHERE id = %d", dto.NewFileName, dto.Id)

	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec(stmt)

	if err != nil {
		return nil, err
	}

	return s.GetFileById(&model.FileGetByIdDTO{Id: dto.Id})
}

func (s *fileStorage) GetFileById(dto *model.FileGetByIdDTO) (*model.FileDTO, error) {
	return s.getFileWithWhereCase(fmt.Sprintf("id = %d", dto.Id))
}
func (s *fileStorage) GetFileByFilename(dto *model.FileGetByNameDTO) (*model.FileDTO, error) {
	return s.getFileWithWhereCase(fmt.Sprintf("filename = '%s' AND folder_id = %d", dto.FileName, dto.FolderId))
}
func (s *fileStorage) GetFilesInFolder(dto *model.FileGetAllDTO) (*[]model.FileDTO, error) {
	return s.getFilesWithWhereCase(fmt.Sprintf("folder_id = %d", dto.FolderId))
}
