package model

type TInterview struct {
	CommonModel
	IDBAdapter[TInterview]
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	InterviewerID uint   `json:"interviewer_id"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	Status        string `json:"status"`
	Feedback      string `json:"feedback"`
}

func (i TInterview) TableName() string {
	return "t_interview"
}
