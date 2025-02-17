package model

type TChat struct {
	CommonModel
	IDBAdapter[TChat] `gorm:"-"`
	SessionID         int64  `json:"sesstion_id"`
	SessionTitle      string `json:"session_title"`
	UserID            int64  `json:"user_id"`
}

func (c TChat) TableName() string {
	return "t_chat"
}
