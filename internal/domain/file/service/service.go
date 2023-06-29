package service

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// "encoding/base64"
	"storage-service/internal/domain/file/model"
	"storage-service/internal/domain/file/storage"
	folder_model "storage-service/internal/domain/folder/model"
)

func saveImg(url, base64Str string) (bool, error) {
	bts, err := base64.StdEncoding.DecodeString(base64Str)

	if err != nil {
		return false, err
	}

	err = ioutil.WriteFile(url, bts, 0644)
	if err != nil {
		return false, err
	}
	return true, nil
}

type fileService struct {
	storage       storage.FileStorage
	folderService FolderService
}

func NewFileService(storage storage.FileStorage, folderfileService FolderService) *fileService {
	return &fileService{
		storage:       storage,
		folderService: folderfileService,
	}
}

func (s *fileService) AddFile(dto *model.FileAddDTO) (*model.FileDTO, error) {
	folderUrl, err := s.folderService.GetUrl(&folder_model.FolderGetByIdDTO{Id: dto.FolderId})
	if folderUrl == "" {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	result, _ := s.storage.GetFileByFilename(&model.FileGetByNameDTO{FileName: dto.FileName, FolderId: dto.FolderId})
	if result != nil {
		return nil, fmt.Errorf("this file in this folder is already exists")
	}

	filename := dto.FileName

	if len(strings.Split(filename, ".")) == 1 || strings.Split(filename, ".")[len(strings.Split(filename, "."))-1] != dto.Format {
		filename += "." + dto.Format
		dto.FileName += "." + dto.Format
	}

	flag, err := saveImg(folderUrl+filename, dto.Base64)
	if err != nil {
		return nil, err
	}
	if !flag {
		return nil, fmt.Errorf("something hapend")
	}
	return s.storage.AddFile(dto)
}
func (s *fileService) DeleteFile(dto *model.FileDeleteDTO) (*model.FileDeleteResponceDTO, error) {
	result, _ := s.GetFileById(&model.FileGetByIdDTO{Id: dto.Id})
	if result == nil {
		return nil, fmt.Errorf("this file is not exist")
	}

	folderUrl, err := s.folderService.GetUrl(&folder_model.FolderGetByIdDTO{Id: result.FolderId})
	if folderUrl == "" {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	err = os.Remove(folderUrl + result.FileName)
	if err != nil {
		return &model.FileDeleteResponceDTO{
			Success: false,
			Mess:    err.Error(),
		}, nil
	}

	return s.storage.DeleteFile(dto)
}
func (s *fileService) RenameFile(dto *model.FileRenameDTO) (*model.FileDTO, error) {
	result, _ := s.GetFileById(&model.FileGetByIdDTO{Id: dto.Id})
	if result == nil {
		return nil, fmt.Errorf("this file is not exist")
	}

	if len(strings.Split(dto.NewFileName, ".")) == 1 || strings.Split(dto.NewFileName, ".")[len(strings.Split(dto.NewFileName, "."))-1] != strings.Split(result.FileName, ".")[1] {
		dto.NewFileName += "." + strings.Split(result.FileName, ".")[1]
	}
	if dto.NewFileName == result.FileName {
		return nil, fmt.Errorf("the new file name is the same as the old file name")
	}

	flag, _ := s.GetFileByFilename(&model.FileGetByNameDTO{FileName: dto.NewFileName, FolderId: result.FolderId})
	if flag != nil {
		return nil, fmt.Errorf("a file with that name already exists")
	}
	folderUrl, err := s.folderService.GetUrl(&folder_model.FolderGetByIdDTO{Id: result.FolderId})
	if folderUrl == "" {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	err = os.Rename(folderUrl+result.FileName, folderUrl+dto.NewFileName)

	if err != nil {
		return nil, err
	}

	return s.storage.RenameFile(dto)
}

func (s *fileService) GetFileById(dto *model.FileGetByIdDTO) (*model.FileDTO, error) {
	return s.storage.GetFileById(dto)
}
func (s *fileService) GetFileByFilename(dto *model.FileGetByNameDTO) (*model.FileDTO, error) {
	return s.storage.GetFileByFilename(dto)
}
func (s *fileService) GetFilesInFolder(dto *model.FileGetAllDTO) (*[]model.FileDTO, error) {
	return s.storage.GetFilesInFolder(dto)
}

func (s *fileService) GetFileBase64(dto *model.FileGetByIdDTO) (*model.FileBase64, error) {
	result, err := s.storage.GetFileById(dto)
	if err != nil {
		return nil, err
	}

	folderUrl, err := s.folderService.GetUrl(&folder_model.FolderGetByIdDTO{Id: result.FolderId})
	if folderUrl == "" {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(folderUrl + result.FileName)
	if err != nil {
		return nil, err
	}

	base64Str := base64.StdEncoding.EncodeToString(data)
	return &model.FileBase64{
		Id:       result.Id,
		FolderId: result.FolderId,
		FileName: result.FileName,
		Base64:   base64Str,
	}, nil
}

func (s *fileService) GetFilesInFolderBase64(dto *model.FileGetAllDTO) (*[]model.FileBase64, error) {
	result, err := s.storage.GetFilesInFolder(dto)
	if err != nil {
		return nil, err
	}

	folderUrl, err := s.folderService.GetUrl(&folder_model.FolderGetByIdDTO{Id: dto.FolderId})
	if folderUrl == "" {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	files := make([]model.FileBase64, len(*result))

	for i, file := range *result {
		data, err := ioutil.ReadFile(folderUrl + file.FileName)
		if err != nil {
			return nil, err
		}
		files[i] = model.FileBase64{
			Id:       file.Id,
			FolderId: file.FolderId,
			FileName: file.FileName,
			Base64:   base64.StdEncoding.EncodeToString(data),
		}
	}

	return &files, nil
}

func (s *fileService) GetFilesInFolderZip(dto *model.FileGetAllDTO) (*model.FilesZip, error) {
	result, err := s.storage.GetFilesInFolder(dto)
	if err != nil {
		return nil, err
	}

	folder, err := s.folderService.GetFolder(&folder_model.FolderGetByIdDTO{Id: dto.FolderId})
	if err != nil {
		return nil, err
	}

	folderUrl, err := s.folderService.GetUrl(&folder_model.FolderGetByIdDTO{Id: dto.FolderId})

	if folderUrl == "" {
		return nil, fmt.Errorf("no such folder")
	}
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, len(*result))
	for i, file := range *result {
		fileNames[i] = file.FileName
	}

	buf := new(bytes.Buffer)
	archive := zip.NewWriter(buf)

	for _, fileName := range fileNames {
		fw, err := archive.Create(fileName)
        if err != nil {
            return nil, err
        }
		data, err := ioutil.ReadFile(folderUrl + fileName)
		if err != nil {
			return nil, err
		}
        _, err = fw.Write(data)
        if err != nil {
            return nil, err
        }
	}
	err = archive.Close()
    if err != nil {
		return nil, err
    }

	if err != nil {
		return nil, err
	}

	return &model.FilesZip{
		FolderId: folder.Id,
		ZipName:  folder.FolderName + ".zip",
		Data:     buf.Bytes(),
	}, nil
}
