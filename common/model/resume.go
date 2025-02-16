package model

import "time"

type TResume struct {
	CommonModel
	IDBAdapter[TResume] `gorm:"-"`
	UserID              int64     `json:"user_id"`   // 代表哪个用户上传的简历
	Url                 string    `json:"url"`       // 存储简历的url
	Path                string    `json:"path"`      // 存储简历的path
	FileName            string    `json:"file_name"` // 文件名
	Size                int64     `json:"size"`      // 文件大小
	UploadTime          time.Time `json:"upload_time"`
	FolderId            int64     `json:"folder_id"` // 文件夹Id
}

func (r TResume) TableName() string {
	return "t_resume"
}
