package ssex

import (
	"Ai-HireSphere/common/zlog"
	"context"
	"fmt"
	"net/http"
	"sync"
)

type Server struct {
	writer http.ResponseWriter
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	data   chan string
	err    error
}

var (
	sseEventDone = "data: [DONE]\n\n"
)

// 升级协议
func Upgrade(ctx context.Context, w http.ResponseWriter) *Server {

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	w.Header().Set("X-Accel-Buffering", "no") // 禁用 Nginx 的缓冲
	w.WriteHeader(http.StatusOK)
	ctx, cancel := context.WithCancel(ctx)

	server := &Server{
		writer: w,
		ctx:    ctx,
		cancel: cancel,
		data:   make(chan string, 10),
	}
	server.Flush()
	server.use()
	return server
}

// 单次写入
func (server *Server) Write(msg string) *Server {

	if server.err != nil {
		return server
	}
	server.wg.Add(1)
	server.data <- msg
	return server
}
func (server *Server) WriteEvent(Event SseEvent) *Server {
	return server.Write(Event.build())
}

// 不断写入
func (server *Server) push() {
	defer server.wg.Done()
	for {
		select {
		case <-server.ctx.Done():
			if len(server.data) != 0 {
				// 这里让他先处理完 如果不处理完会会wg.wait()一直阻塞
				continue
			}
			zlog.InfofCtx(server.ctx, "stop push")
			return
		case msg := <-server.data:
			server.wg.Done()
			if server.err != nil {
				continue
			}
			_, server.err = fmt.Fprint(server.writer, msg)
			if server.err != nil {
				zlog.ErrorfCtx(server.ctx, "push failed, err: %v", server.err)
				return
			}
			server.Flush()
		}
	}
}

// 使用sse
func (server *Server) use() {
	server.wg.Add(1)
	zlog.DebugfCtx(server.ctx, "use sse")
	go server.push()
}

// 关闭sse
func (server *Server) Close() *Server {
	server.cancel()
	server.wg.Wait()
	fmt.Fprint(server.writer, sseEventDone)
	server.err = fmt.Errorf("close sse")
	close(server.data)
	return server
}

func (server *Server) Flush() *Server {
	if flusher, ok := server.writer.(http.Flusher); ok && flusher != nil {
		flusher.Flush()
	}
	return server
}

func (server *Server) Error() error {
	return server.err
}
