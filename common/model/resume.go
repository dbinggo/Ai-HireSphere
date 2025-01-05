package model

type TResume struct {
	CommonModel
	IDBAdapter[TResume]
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	Content    string `json:"content"` // JSON格式存储简历内容
	UploadTime string `json:"upload_time"`
}

func (r TResume) TableName() string {
	return "t_resume"
}
