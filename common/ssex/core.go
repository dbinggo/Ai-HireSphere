package ssex

import (
	"Ai-HireSphere/common/zlog"
	"context"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
)

type Server struct {
	writer http.ResponseWriter

	ctx    context.Context
	cancel context.CancelFunc

	data chan interface{}
	err  error
}

// 升级协议
func Upgrade(ctx context.Context, w http.ResponseWriter) Server {
	//设置响应头
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache, no-transform")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no") // 禁用 Nginx 的缓冲

	ctx, cancel := context.WithCancel(ctx)

	server := Server{
		writer: w,

		ctx:    ctx,
		cancel: cancel,
		data:   make(chan interface{}),
	}
	server.use()
	return server
}

// 单次写入
func (server *Server) Write(resp interface{}) error {

	if server.err != nil {
		close(server.data)
		return server.err
	}
	server.data <- resp

	return nil
}

// 不断写入
func (server *Server) push() {

	var data []byte
	me := context.WithoutCancel(server.ctx)
	for {
		select {
		case <-me.Done():
			zlog.InfofCtx(server.ctx, "stop push")
			return
		case resp := <-server.data:
			data, server.err = json.Marshal(resp)
			if server.err != nil {
				panic(server.err)
			}
			_, server.err = server.writer.Write(append(data, []byte("\n\n")...))
			if server.err != nil {
				zlog.ErrorfCtx(server.ctx, "push failed, err: %v", server.err)
			}
			server.Flush()
		}
	}
}

// 使用sse
func (server *Server) use() {
	zlog.DebugfCtx(server.ctx, "use sse")
	go server.push()
}

func (server *Server) Close() {

	server.Flush()
	server.cancel()
}

func (server *Server) Flush() {
	if f, ok := server.writer.(http.Flusher); ok {
		f.Flush()
	}
}
