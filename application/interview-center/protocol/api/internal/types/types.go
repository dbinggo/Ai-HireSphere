// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5

package types

type Resume struct {
	UserID int64  `json:"user_id" form:"user_id"`
	Name   string `json:"name" form:"name"`
}

type ResumeAddRequest struct {
	Resume
}

type ResumeAddResponse struct {
	ID int64 `json:"id"`
}
