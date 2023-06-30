package model

import "fmt"

type FileDTO struct {
	Id       int64
	FolderId int64
	FileName string
	Url      string
}

type FileBase64 struct {
	Id       int64
	FolderId int64
	FileName string
	Base64   string
}

type FilesZip struct {
	ZipName string
	Url     string
}

type FileAddDTO struct {
	Base64   string
	FolderId int64
	FileName string
	Format   string
	Url      string
}

func (dto *FileAddDTO) ExtractInsertSQL() string {
	return fmt.Sprintf(` (%d, '%s')`, dto.FolderId, dto.FileName)
}

type FileDeleteDTO struct {
	Id int64
}

type FileDeleteResponceDTO struct {
	Success bool
	Mess    string
}

type FileRenameDTO struct {
	Id          int64
	NewFileName string
}

type FileGetByIdDTO struct {
	Id int64
}

type FileGetByNameDTO struct {
	FileName string
	FolderId int64
}

type FileGetAllDTO struct {
	FolderId int64
	// UserId   int64
}
