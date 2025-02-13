package interview

import (
	"Ai-HireSphere/common/ssex"
	"Ai-HireSphere/common/thrift/deepseek"
	"context"
	"github.com/zeromicro/go-zero/core/trace"
	"net/http"
	"time"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.SSEReq, w http.ResponseWriter) error {

	// 协议升级
	sse := ssex.Upgrade(l.ctx, w)

	// 模拟思考过程
	repet := "他刚刚说了：" + req.Data

	for _, rep := range repet {
		sse.WriteEvent(ssex.SseEvent{
			Data: deepseek.NormalResp{
				Id:    trace.TraceIDFromContext(l.ctx),
				Model: "代金鱼-R1",
				Choices: []deepseek.NormalRespChoices{
					{
						Index: 0,
						Message: deepseek.NormalRespMessage{
							Role:             "assistant",
							ReasoningContent: string(rep),
							Content:          "",
						},
						FinishReason: "",
					},
				},
				Usage:             deepseek.NormalRespUsage{},
				SystemFingerprint: "",
			},
		})
		time.Sleep(time.Second)
	}
	ans := "您刚刚说了:" + req.Data
	// 模拟回答过程
	for _, rep := range ans {
		sse.WriteEvent(ssex.SseEvent{
			Data: deepseek.NormalResp{
				Id:    trace.TraceIDFromContext(l.ctx),
				Model: "代金鱼-R1",
				Choices: []deepseek.NormalRespChoices{
					{
						Index: 0,
						Message: deepseek.NormalRespMessage{
							Role:             "assistant",
							ReasoningContent: "",
							Content:          string(rep),
						},
						FinishReason: "",
					},
				},
				Usage:             deepseek.NormalRespUsage{},
				SystemFingerprint: "",
			},
		})
		time.Sleep(time.Second)
	}

	defer sse.Close()

	return nil
}
