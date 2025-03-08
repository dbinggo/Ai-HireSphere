// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	interview "Ai-HireSphere/application/interview/interfaces/api/internal/handler/interview"
	resume "Ai-HireSphere/application/interview/interfaces/api/internal/handler/resume"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					// 进行单次对话
					Method:  http.MethodPost,
					Path:    "/chat",
					Handler: interview.ChatHandler(serverCtx),
				},
				{
					// 面试对话
					Method:  http.MethodPost,
					Path:    "/chat/agent",
					Handler: interview.ChatInterviewHandler(serverCtx),
				},
				{
					// 生成会话 id
					Method:  http.MethodPost,
					Path:    "/chat/new",
					Handler: interview.ChatNewHandler(serverCtx),
				},
				{
					// 新建一场面试
					Method:  http.MethodPost,
					Path:    "/new_interview",
					Handler: interview.CreateInterviewHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1/interview"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					// 筛选简历
					Method:  http.MethodPost,
					Path:    "/check",
					Handler: resume.CheckResumeHandler(serverCtx),
				},
				{
					// 删除简历
					Method:  http.MethodDelete,
					Path:    "/delete/:id",
					Handler: resume.DeleteResumeHandler(serverCtx),
				},
				{
					// 简历评估
					Method:  http.MethodPost,
					Path:    "/evaluate",
					Handler: resume.EvaluateResumeHandler(serverCtx),
				},
				{
					// 新建简历文件夹
					Method:  http.MethodPost,
					Path:    "/folder",
					Handler: resume.CreateResumeFolderHandler(serverCtx),
				},
				{
					// 更新简历文件夹
					Method:  http.MethodPut,
					Path:    "/folder",
					Handler: resume.UpdateResumeFolderHandler(serverCtx),
				},
				{
					// 删除简历文件夹
					Method:  http.MethodDelete,
					Path:    "/folder/:id",
					Handler: resume.DeleteResumeFolderHandler(serverCtx),
				},
				{
					// 获取简历文件夹列表
					Method:  http.MethodGet,
					Path:    "/folder/list",
					Handler: resume.GetResumeFolderListHandler(serverCtx),
				},
				{
					// 获取简历
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: resume.GetResumeListHandler(serverCtx),
				},
				{
					// 上传简历
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: resume.UploadResumeHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1/resume"),
	)
}
