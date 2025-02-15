// Code generated by goctl. DO NOT EDIT.
package types

type CommonListReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

type CommonListResp struct {
	Total int64 `json:"total"`
}

type DeleteResumeReq struct {
	ResumeId int64 `path:"id"`
}

type GetResumeListReq struct {
	CommonListReq
}

type GetResumeListResp struct {
	CommonListResp
	List []ResumeInfo `json:"list"`
}

type ResumeInfo struct {
	ResumeId   int64  `json:"id"`
	ResumeName string `json:"name"`
	ResumeUrl  string `json:"url"`
	UploadTime string `json:"upload_time"`
	ResumeSize int64  `json:"size"`
	UserId     int64  `json:"user_id"`
}

type SSEReq struct {
	Data string `json:"data"`
}

type UploadResumeResp struct {
	Address string `json:"address"`
}
