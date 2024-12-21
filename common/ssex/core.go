package ssex

import (
	"context"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"time"
)

type Server struct {
	writer       http.ResponseWriter
	heartTimeout time.Duration
	log          Log

	ctx    context.Context
	cancel context.CancelFunc

	data chan interface{}
	err  error
}

// 升级协议
func Upgrade(w http.ResponseWriter, log Log) Server {
	//设置响应头
	header := w.Header()
	header.Set("Content-Type", "text/event-stream")
	header.Set("Cache-Control", "no-cache")
	header.Set("Connection", "keep-alive")

	ctx, cancel := context.WithCancel(context.Background())

	timeout := time.Second * 3

	server := Server{
		writer:       w,
		heartTimeout: timeout,
		log:          log,
		ctx:          ctx,
		cancel:       cancel,
		data:         make(chan interface{}),
	}
	server.use()
	return server
}

// 单次写入
func (server *Server) Write(resp interface{}) error {

	server.data <- resp

	if server.err != nil {
		close(server.data)
		return server.err
	}

	return nil
}

// 不断写入
func (server *Server) push() {
	var err error
	defer func() {
		if err != nil {
			server.err = err
		}
	}()

	var data []byte
	me := context.WithoutCancel(server.ctx)
	for {
		select {
		case <-me.Done():
			server.log.Info("stop push")
			return
		case resp := <-server.data:
			err = json.Unmarshal(data, resp)
			if err != nil {
				panic(err)
			}
			_, err = server.writer.Write(data)
			if err != nil {
				server.log.Errorf("push failed, err: %v", err)
			}
		}
	}
}

// 心跳
func (server *Server) heart() {
	var err error
	defer func() {
		if err != nil {
			server.err = err
		}
	}()

	var info string
	info = "hello"

	me := context.WithoutCancel(server.ctx)
	for {
		select {
		case <-me.Done():
			return
		default:
			_, err = server.writer.Write([]byte(info))
			if err != nil {
				server.cancel()
				server.log.Warn("client is dead")
				return
			}
			time.Sleep(server.heartTimeout)
		}

	}
}

// 使用sse
func (server *Server) use() {
	go server.heart()
	go server.push()
}
