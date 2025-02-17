package model

type TFolder struct {
	CommonModel
	IDBAdapter[TFolder] `gorm:"-"`
	UserId              int64  `json:"user_id" gorm:"column:user_id;type:bigint;not null;comment:用户id"`
	Name                string `json:"name" gorm:"column:name;type:varchar(64);not null;comment:文件夹名称"`
}

func (r TFolder) TableName() string {
	return "t_folder"
}
