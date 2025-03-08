package codex

import (
	"context"
	"fmt"
	"strconv"

	"Ai-HireSphere/common/codex/types"

	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// 通过断言，强制编译器确认XCode接口是否完全被Status实现
var _ CodeX = (*Status)(nil)

// 这里是为了能传入proto流中
type Status struct {
	sts *types.Status
}

func Error(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

func Errorf(code Code, format string, args ...interface{}) *Status {
	code.msg = fmt.Sprintf(format, args...)
	return Error(code)
}

func (s *Status) Error() string {
	return s.Message()
}

func (s *Status) Code() int {
	return int(s.sts.Code)
}

func (s *Status) Message() string {
	if s.sts.Message == "" {
		return strconv.Itoa(int(s.sts.Code))
	}

	return s.sts.Message
}

// 原理是，在protobuf中的detail中塞入自定义状态code和message（怎么塞：请看status.proto），现在就是从detail中提取数据
func (s *Status) Details() []interface{} {
	if s == nil || s.sts == nil {
		return nil
	}
	details := make([]interface{}, 0, len(s.sts.Details))
	for _, d := range s.sts.Details {
		detail := &ptypes.DynamicAny{}

		if err := d.UnmarshalTo(detail); err != nil {
			details = append(details, err)
			continue
		}
		details = append(details, detail.Message)
	}

	return details
}

// 将自定义message塞入detail中
func (s *Status) WithDetails(msgs ...proto.Message) (*Status, error) {
	for _, msg := range msgs {
		anyMsg, err := anypb.New(msg)
		if err != nil {
			return s, err
		}
		s.sts.Details = append(s.sts.Details, anyMsg)
	}

	return s, nil
}

// 将status转化为proto能识别的
func (s *Status) Proto() *types.Status {
	return s.sts
}

// 将code转换为status
func FromCode(code Code) *Status {
	return &Status{sts: &types.Status{Code: int32(code.Code()), Message: code.Message()}}
}

// 这里开始把不确定的msg类型转化成自定义code和message（也就是转化为status啦），进一步提取数据
func FromProto(pbMsg proto.Message) CodeX {
	msg, ok := pbMsg.(*types.Status)
	if ok {
		if len(msg.Message) == 0 || msg.Message == strconv.FormatInt(int64(msg.Code), 10) {
			return Code{code: int(msg.Code)}
		}
		return &Status{sts: msg}
	}

	return Errorf(ServerErr, "invalid proto message get %v", pbMsg)
}

// 把status转换为code
func toXCode(grpcStatus *status.Status) Code {
	grpcCode := grpcStatus.Code()
	switch grpcCode {
	case codes.OK:
		return OK
	case codes.InvalidArgument:
		return RequestErr
	case codes.NotFound:
		return NotFound
	case codes.PermissionDenied:
		return AccessDenied
	case codes.Unauthenticated:
		return Unauthorized
	case codes.ResourceExhausted:
		return LimitExceed
	case codes.Unimplemented:
		return MethodNotAllowed
	case codes.DeadlineExceeded:
		return Deadline
	case codes.Unavailable:
		return ServiceUnavailable
	case codes.Unknown:
		return String(grpcStatus.Message())
	}

	return ServerErr
}

// error转化为xcode
func CodeFromError(err error) CodeX {
	err = errors.Cause(err)
	if code, ok := err.(CodeX); ok {
		return code
	}

	switch err {
	case context.Canceled:
		return Canceled
	case context.DeadlineExceeded:
		return Deadline
	}

	return add(500, err.Error())
}

// error转换为status
func FromError(err error) *status.Status {
	err = errors.Cause(err)
	if code, ok := err.(CodeX); ok {
		grpcStatus, e := gRPCStatusFromXCode(code)
		if e == nil {
			return grpcStatus
		}
	}

	var grpcStatus *status.Status
	switch err {
	case context.Canceled:
		grpcStatus, _ = gRPCStatusFromXCode(Canceled)
	case context.DeadlineExceeded:
		grpcStatus, _ = gRPCStatusFromXCode(Deadline)
	default:
		grpcStatus, _ = status.FromError(err)
	}

	return grpcStatus
}

func gRPCStatusFromXCode(code CodeX) (*status.Status, error) {
	var sts *Status
	switch v := code.(type) {
	case *Status:
		sts = v
	case Code:
		sts = FromCode(v)
	default:
		sts = Error(Code{code.Code(), code.Message()})
		for _, detail := range code.Details() {
			if msg, ok := detail.(proto.Message); ok {
				_, _ = sts.WithDetails(msg)
			}
		}
	}

	stas := status.New(codes.Unknown, strconv.Itoa(sts.Code()))
	return stas.WithDetails(sts.Proto())
}

// 把status转化回xcode
func GrpcStatusToXCode(gstatus *status.Status) CodeX {
	//这里初步提取detail中的数据
	details := gstatus.Details()
	for i := len(details) - 1; i >= 0; i-- {
		detail := details[i]
		if pb, ok := detail.(proto.Message); ok {
			//进一步提取，拿到最终信息
			return FromProto(pb)
		}
	}

	return toXCode(gstatus)
}
