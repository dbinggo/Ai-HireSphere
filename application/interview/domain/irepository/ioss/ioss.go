package ioss

import "io"

type Ioss interface {
	UploadFile(objectName string, reader io.Reader) error
	DownloadFile(objectName, filePath string) error
	DeleteFile(objectName string) error
}
