package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

type FileBase struct {
	FileByte []byte
	Filename string
	Ext      string
}

func ReadFormFile(r *http.Request, key string) FileBase {
	_, header, err := r.FormFile(key)
	if err != nil {
		return FileBase{}
	}
	base, err := getFileBase(header)
	if err != nil {
		return FileBase{}
	}

	return base
}

func ReadFormFiles(r *http.Request, key string) []FileBase {
	var bases []FileBase
	headers := r.MultipartForm.File[key]
	for i := 0; i < len(headers); i++ {
		base, err := getFileBase(headers[i])
		if err != nil {
			return nil
		}
		bases = append(bases, base)
	}

	return bases
}

func getFileBase(header *multipart.FileHeader) (FileBase, error) {
	var base FileBase
	var err error
	var file multipart.File
	file, err = header.Open()
	if err != nil {
		return FileBase{}, err
	}

	base.FileByte, err = io.ReadAll(file)
	if err != nil {
		return FileBase{}, err
	}

	ext := filepath.Ext(header.Filename)
	base.Ext = strings.Split(ext, ".")[0]
	base.Filename = strings.Split(header.Filename, ext)[0]
	return base, nil
}
