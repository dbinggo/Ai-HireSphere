package app

import (
	"Ai-HireSphere/application/interview/domain/irepository"
	"Ai-HireSphere/application/interview/domain/irepository/ioss"
	"Ai-HireSphere/application/interview/domain/model/entity"
	"Ai-HireSphere/application/interview/domain/service"
	"Ai-HireSphere/common/coze"
	"context"
	"github.com/dbinggo/gerr"
	"mime/multipart"
)

type IResumeApp interface {
	UploadResume(ctx context.Context, file multipart.File, handler *multipart.FileHeader) (string, gerr.Error)
	ListResume(ctx context.Context, userId int64, page int, pageSize int) (int64, []entity.ResumeEntity, gerr.Error)
	DeleteResume(ctx context.Context, id int64) gerr.Error
	CreateSession(ctx context.Context, userID int64) (int64, gerr.Error)
	Chat(ctx context.Context, id int64, message string) (chan coze.BotStreamReply, gerr.Error)
	UpdateChatName(ctx context.Context, id int64, name string) gerr.Error
	ListChats(ctx context.Context, userID int64, pageSize, pageNum int) (int64, []entity.ChatEntity, gerr.Error)
	GetChatHistory(ctx context.Context, id int64) ([]coze.BotMessage, gerr.Error)
}

type ResumeApp struct {
	Oss     ioss.Ioss
	CozeApi *coze.CozeApi
	Repo    irepository.IRepoBroker
}

func (r *ResumeApp) CreateSession(ctx context.Context, userID int64) (int64, gerr.Error) {
	return service.NewChatService(context.Background(), r.Repo, *r.CozeApi).CreateSession(userID)
}

func (r *ResumeApp) Chat(ctx context.Context, id int64, message string) (chan coze.BotStreamReply, gerr.Error) {
	return service.NewChatService(ctx, r.Repo, *r.CozeApi).Chat(id, message)
}

func (r *ResumeApp) UpdateChatName(ctx context.Context, id int64, name string) gerr.Error {
	return service.NewChatService(ctx, r.Repo, *r.CozeApi).UpdateChatName(id, name)
}

func (r *ResumeApp) ListChats(ctx context.Context, userID int64, pageSize, pageNum int) (int64, []entity.ChatEntity, gerr.Error) {
	return service.NewChatService(ctx, r.Repo, *r.CozeApi).ListChats(userID, pageSize, pageNum)
}

func (r *ResumeApp) GetChatHistory(ctx context.Context, id int64) ([]coze.BotMessage, gerr.Error) {
	return service.NewChatService(ctx, r.Repo, *r.CozeApi).GetChatHistory(id)
}

func NewResumeApp(oss ioss.Ioss, repo irepository.IRepoBroker, api *coze.CozeApi) *ResumeApp {
	return &ResumeApp{
		Oss:     oss,
		CozeApi: api,
		Repo:    repo,
	}

}

func (r *ResumeApp) UploadResume(ctx context.Context, file multipart.File, handler *multipart.FileHeader) (string, gerr.Error) {
	// 调用服务
	resume, err := service.NewResumeService(ctx, r.Oss, r.Repo).UploadResume(file, handler)
	if err != nil {
		return "", err
	}
	return resume.Url, nil
}

func (r *ResumeApp) ListResume(ctx context.Context, userId int64, page int, pageSize int) (int64, []entity.ResumeEntity, gerr.Error) {
	count, resp, err := r.Repo.ListResume(ctx, userId, page, pageSize)
	if err != nil {
		return 0, nil, err
	}
	return count, resp, nil
}

func (r *ResumeApp) DeleteResume(ctx context.Context, id int64) gerr.Error {
	err := service.NewResumeService(ctx, r.Oss, r.Repo).DeleteResume(id)
	return err
}
