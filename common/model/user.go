package model

type TUser struct {
	CommonModel
	IDBAdapter[TUser] `gorm:"-"` // 通过继承的方式，将通用的数据库操作方法注入到User中
	Username          string     `json:"username"`
	Email             string     `json:"email" gorm:"type:varchar(255)"`
	Avatar            string     `json:"avatar" gorm:"type:varchar(255)"`
	Phone             string     `json:"phone" gorm:"uniqueIndex;type:varchar(64)"`
	Role              string     `json:"role" gorm:"type:varchar(64)"` // 新增字段：用户角色，例如管理员、面试官等
	Sex               int        `json:"sex" gorm:"type:int"`
}

func (u TUser) TableName() string {
	return "t_user"
}
