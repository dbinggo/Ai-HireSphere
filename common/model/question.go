package model

type TQuestion struct {
	CommonModel
	IDBAdapter[TQuestion]
	ID      uint   `json:"id"`
	Content string `json:"content"`
	Type    string `json:"type"`
	Options string `json:"options"` // JSON格式存储选项
	Answer  string `json:"answer"`
}

func (q TQuestion) TableName() string {
	return "t_question"
}
