package repository

import (
	"Ai-HireSphere/common/utils"
)

// 简历实体
type Resume struct {
	ID     uint           `json:"id" redis:"id"`
	UserID uint           `json:"user_id" redis:"user_id"`
	Name   string         `json:"name" redis:"name"`
	File   utils.FileBase `json:"file"`
}
