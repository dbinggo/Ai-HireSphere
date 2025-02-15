package model

type TQuestion struct {
	CommonModel
	IDBAdapter[TQuestion] `gorm:"-"`
	Content               string `json:"content"` // 问题
	Answer                string `json:"answer"`  // 答案
	Url                   string `json:"url"`     // 问题来源
	Company               string `json:"company"` // 面试公司
	Type                  int    `json:"type"`    // 问题类型 使用bit放置位置(前端 后端 运维 测试 八股 计网 计算机系统 数据库 简历 场景题 算法 java go c++)
}

func (q TQuestion) TableName() string {
	return "t_question"
}
