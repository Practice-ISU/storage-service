package model

type FileDTO struct {
	Id       int64
	FolderId int64
	FileName string
	Url      string
}

type FileAddDTO struct {
	FolderId int64
	FileName string
}

type FileDeleteDTO struct {
	Id int64
}

type FileDeleteResponceDTO struct {
	Success bool
	Mess string
}

type FileRenameDTO struct {
	Id          int64
	NewFileName string
}

type FileGetDTO struct {
	Id int64
}

type FileGetAllDTO struct {
	Id int64
}

