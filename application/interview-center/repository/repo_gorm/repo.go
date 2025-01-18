package repo_gorm

import (
	"Ai-HireSphere/common/coze"
	"gorm.io/gorm"
)

type ResumeModel struct {
	gorm.Model
	UserID uint   `json:"user_id" redis:"user_id"`
	Name   string `json:"name" redis:"name"`
}

type Repo struct {
	Model ResumeModel
	DB    *gorm.DB
	Coze  *coze.CozeApi
}

func (b ResumeModel) TableName() string {
	return "resumes"
}
