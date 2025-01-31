package entity

import (
	"Ai-HireSphere/application/interview/domain/irepository/ioss"
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/model"
	"Ai-HireSphere/common/zlog"
	"fmt"
	"github.com/dbinggo/gerr"
	"mime/multipart"
	"time"
)

// 简历实体
type ResumeEntity struct {
	Id         int64
	Url        string
	UserId     int64
	UploadTime time.Time
	Size       int64
	Path       string
	FileName   string

	Handler *multipart.FileHeader
	File    multipart.File
}

// 实体与schema互转
var _ model.ICommonEntity[*ResumeEntity, *model.TResume] = &ResumeEntity{}

// 实现

func (r *ResumeEntity) Transform() *model.TResume {
	return &model.TResume{
		CommonModel: model.CommonModel{
			ID: r.Id,
		},
		IDBAdapter: &model.CommonAdapter[model.TResume]{},
		Url:        r.Url,
		UserID:     r.UserId,
		Size:       r.Handler.Size,
		UploadTime: r.UploadTime,
	}

}

func (r *ResumeEntity) From(f *model.TResume) *ResumeEntity {
	r.Id = f.ID
	r.UserId = f.UserID
	r.Url = f.Url
	r.UploadTime = f.UploadTime
	r.FileName = f.FileName
	r.Size = f.Size
	return r
}

// ValidateUpload
//
//	@Description: 校验是否满足上传条件 上传条件:有file流，有句炳，有文件名，有上传的用户,有大小
//	@receiver r
//	@return gerr.Error
func (r *ResumeEntity) ValidateUpload() gerr.Error {
	if r.Handler == nil {
		return gerr.Wraps(codex.ResumeUploadEmpty)
	}
	const MAX_RESUME_SIZE = 1024 * 1024 * 20 // 20M

	if r.Handler.Size > MAX_RESUME_SIZE {
		return gerr.Wraps(codex.ResumeUploadMAX)
	}
	if r.File == nil {
		return gerr.Wraps(codex.ResumeUploadEmpty)
	}
	if r.UserId == 0 {
		return gerr.Wraps(codex.ResumeUploadFail)
	}
	if r.FileName == "" {
		return gerr.Wraps(codex.ResumeUploadEmpty)
	}
	if r.Handler.Size == 0 {
		return gerr.Wraps(codex.ResumeUploadEmpty)
	}
	return nil
}

func (r *ResumeEntity) UploadResume(oss ioss.Ioss) gerr.Error {
	if err := r.ValidateUpload(); err != nil {
		return err
	}
	r.GeneratePathAndUrl()
	objectName := r.Path
	err := oss.UploadFile(objectName, r.File)
	if err != nil {
		zlog.Errorf("upload file fail: %v", err)
		return gerr.Wraps(codex.ResumeUploadFail, err)
	}
	// 上传完了
	zlog.Debugf("简历%s上传完成,路径 %s ", r.FileName, r.Path)
	return nil
}

func (r *ResumeEntity) GeneratePathAndUrl() {
	r.Path = fmt.Sprintf("resume/%d/%d_%s", r.UserId, time.Now().Unix(), r.FileName)
	r.Url = fmt.Sprintf("https://%s%s", "ai-hiresphere.oss-cn-beijing.aliyuncs.com/", r.Path)
}
