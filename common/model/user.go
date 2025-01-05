package model

type TUser struct {
	CommonModel
	IDBAdapter[TUser]        // 通过继承的方式，将通用的数据库操作方法注入到User中
	Username          string `json:"username"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Role              string `json:"role"` // 新增字段：用户角色，例如管理员、面试官等
	Sex               int    `json:"sex"`
}

func (u TUser) TableName() string {
	return "t_user"
}
