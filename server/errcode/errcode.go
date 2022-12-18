package errcode

import "google.golang.org/grpc/status"

var (
	ErrLogicParam = status.Error(40000, "请求参数异常，请检查参数")

	ErrRpcRequestFail = status.Error(50000, "系统请求错误，请稍后重试")
)
