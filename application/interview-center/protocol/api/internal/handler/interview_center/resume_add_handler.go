package interview_center

import (
	"fmt"
	"net/http"
	"time"

	"Ai-HireSphere/application/interview-center/protocol/api/internal/logic/interview_center"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/svc"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ResumeAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResumeAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interview_center.NewResumeAddLogic(r.Context(), svcCtx)
		resp, err := l.ResumeAdd(&req)
		// 设置响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// 启用长连接
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming is unsupported!", http.StatusInternalServerError)
			return
		}

		// 发送事件
		for {
			// 发送当前时间
			fmt.Fprintf(w, "data: %s\n\n", time.Now().String())
			flusher.Flush() // 确保将数据刷新到客户端

			// 发送心跳消息（注释事件）
			fmt.Fprintf(w, ": keep-alive\n\n")
			flusher.Flush() // 刷新心跳消息

			// 每隔 1 秒发送一次数据
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
