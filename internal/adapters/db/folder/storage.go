package storage

import (
	"fmt"
	"storage-service/internal/domain/folder/model"
	"storage-service/pkg/clients/postgre"
)

type folderStorage struct {
	psqlClient *postgre.PsqlClient
}

func NewFileStorage(cnf postgre.Config) *folderStorage {
	client := postgre.NewPsqlCliennt(cnf)
	return &folderStorage{
		psqlClient: client,
	}
}

type FolderScanner interface {
	Scan(dest ...any) error
}

func scanFolder(scanner FolderScanner) (*model.FolderDTO, error) {
	folder := &model.FolderDTO{}
	err := scanner.Scan(&folder.Id, &folder.UserId, &folder.FolderName)
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (s *folderStorage) getFolderWithWhereCase(whereCase string) (*model.FolderDTO, error) {
	stmt := "SELECT id, user_id, folder_name FROM folders"
	if whereCase != "" {
		stmt += " WHERE " + whereCase
	}
	stmt += " LIMIT 1"

	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result := db.QueryRow(stmt)
	if err != nil {
		return nil, err
	}
	return scanFolder(result)
}

func (s *folderStorage) AddFolder(dto *model.FolderAddDTO) (*model.FolderDTO, error) {
	stmt := fmt.Sprintf("INSERT INTO folders (user_id, folder_name) VALUES %s", dto.ExtractInsertSQL())
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
	return &model.FolderDTO{
		Id:         id,
		UserId:     dto.UserId,
		FolderName: dto.FolderName,
	}, nil
}
func (s *folderStorage) DeleteFolder(dto *model.FolderDeleteDTO) (*model.FolderDeleteResponse, error) {
	stmt := fmt.Sprintf("DELETE FROM folders WHERE id = %d LIMIT 1", dto.Id)
	db, err := s.psqlClient.GetDb()
	if err != nil {
		return &model.FolderDeleteResponse{
			Success: false,
			Mess:    err.Error(),
		}, nil
	}
	defer db.Close()

	_, err = db.Exec(stmt)

	if err != nil {
		return &model.FolderDeleteResponse{
			Success: false,
			Mess:    err.Error(),
		}, nil
	}
	return &model.FolderDeleteResponse{
		Success: true,
		Mess:    "",
	}, nil
}
func (s *folderStorage) RenameFolder(dto *model.FolderRenameDTO) (*model.FolderDTO, error) {
	stmt := fmt.Sprintf("UPDATE folders SET folder_name = %s'' WHERE id = %d LIMIT 1", dto.NewName, dto.Id)
	db, err := s.psqlClient.GetDb()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec(stmt + "RETURNING id;")

	if err != nil {
		return nil, err
	}
	// return &model.FolderDTO{

	// }
	return nil, nil
}

func (s *folderStorage) GetFolderById(dto *model.FolderGetDTO) (*model.FolderDTO, error) {
	return s.getFolderWithWhereCase(fmt.Sprintf(`id = %d AND user_id = %d`, dto.Id, dto.UserId))
}
func (s *folderStorage) GetFolderByName(dto *model.FolderGetDTO) (*model.FolderDTO, error) {
	return s.getFolderWithWhereCase(fmt.Sprintf(`folder_name = '%s' AND user_id = %d`, dto.FolderName, dto.UserId))
}
func (s *folderStorage) GetAllFolders(userId int64) (*[]model.FolderDTO, error) {
	return nil, nil
}
