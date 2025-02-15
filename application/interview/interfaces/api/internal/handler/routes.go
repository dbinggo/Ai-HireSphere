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
			}...,
		),
		rest.WithPrefix("/v1/interview"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					// 删除简历
					Method:  http.MethodDelete,
					Path:    "/delete/:id",
					Handler: resume.DeleteResumeHandler(serverCtx),
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
