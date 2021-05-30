package xcontext

import (
	"context"
	"time"
)

// EmptyContext
type EmptyContext struct {
	RawContext context.Context
	Protocol   ProtocolType
	Trace      RasseTrace
	Mesh       MeshTrace
	User       UserInfo
	Biz        BizTrace
	Lang       LangInfo
}

// 返回一个空的上下文
func NewEmptyContext() *EmptyContext {
	var ctx = &EmptyContext{
		RawContext: context.Background(),
	}
	return ctx
}

// ===== rasseContext start ======
func (ctx *EmptyContext) GetRawContext() interface{} {
	return nil
}

func (ctx *EmptyContext) GetRasseTrace() RasseTrace {
	return ctx.Trace
}

func (ctx *EmptyContext) GetMeshTrace() MeshTrace {
	return ctx.Mesh
}

func (ctx *EmptyContext) GetProtocol() ProtocolType {
	return ctx.Protocol
}

func (ctx *EmptyContext) GetUserInfo() UserInfo {
	return ctx.User
}

func (ctx *EmptyContext) GetBizTrace() BizTrace {
	return ctx.Biz
}

func (ctx *EmptyContext) GetLang() LangInfo {
	return ctx.Lang
}

func (ctx *EmptyContext) Get(key string) (value interface{}, exist bool) {

	return nil, false
}

func (ctx *EmptyContext) Set(key string, value interface{}) {

}

func (ctx *EmptyContext) Deadline() (deadline time.Time, ok bool) {
	if ctx.RawContext != nil {
		deadline, ok = ctx.RawContext.Deadline()
	}
	return
}

func (ctx *EmptyContext) Done() <-chan struct{} {
	if ctx.RawContext != nil {
		return ctx.RawContext.Done()
	}
	return nil
}

func (ctx *EmptyContext) Err() error {
	if ctx.RawContext != nil {
		return ctx.RawContext.Err()
	}
	return nil
}

func (ctx *EmptyContext) Value(key interface{}) interface{} {
	if ctx.RawContext != nil {
		return ctx.RawContext.Value(key)
	}
	return nil
}
