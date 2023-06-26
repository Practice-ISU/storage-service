package model

import "fmt"

type FolderDTO struct {
	Id         int64
	UserId     int64
	FolderName string
}

type FolderAddDTO struct {
	UserId     int64
	FolderName string
}

func (dto *FolderAddDTO) ExtractInsertSQL() string {
	return fmt.Sprintf(" (%d, '%s')", dto.UserId, dto.FolderName)
}

type FolderDeleteDTO struct {
	Id int64
	UserId int64
}

func (dto *FolderDeleteDTO) ExtractDeleteSQL() string {
	return fmt.Sprintf("id = %d", dto.Id)
}

type FolderDeleteResponse struct {
	Success bool
	Mess string
}

type FolderRenameDTO struct {
	Id int64
	NewName string
	UserId int64
}

func (dto *FolderRenameDTO) ExtractRenameSQL() string {
	return fmt.Sprintf("id = %d", dto.Id)
}

type FolderGetDTO struct {
	Id int64
	FolderName string
	UserId int64
}
