package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/heyu52/Fly/i18n"
	"github.com/heyu52/Fly/xcontext"
	"github.com/heyu52/Fly/xlog"
	"github.com/heyu52/Fly/xstring"
	"net/http"
	"time"
)

type xGinContext struct {
	protocol  xcontext.ProtocolType
	flyTrace  xcontext.FlyTrace
	meshTrace xcontext.MeshTrace
	user      xcontext.UserInfo
	lang      xcontext.LangInfo
	bizTrace  xcontext.BizTrace
	raw       *gin.Context
}

func NewContex(ginContext *gin.Context) xcontext.XContext {
	var ctx xGinContext
	ctx.raw = ginContext
	ctx.protocol = xcontext.Gin
	ctx.initUserInfo()
	ctx.initMeshInfo()
	ctx.initFlyTraceInfo()
	ctx.initLang()
	ctx.initBizTraceInfo()
	return &ctx
}

func (ctx *xGinContext) initUserInfo() {
	var info xcontext.UserInfo
	var staff = GetCurrentUser(ctx.raw)
	if staff != nil {
		info.UserId = staff.StaffID
		info.UserName = staff.StaffName
	}

	ctx.user=info
}


// private Init
// 初始化链路跟踪信息
func (ctx *xGinContext) initMeshInfo() {
	var info = make(xcontext.MeshTrace)
	headerKeys := xcontext.GetMeshKeys()
	for _, v := range headerKeys {
		value := ctx.raw.GetHeader(v)
		if value != "" {
			info[v] = value
		}
	}
	ctx.meshTrace = info
}


func (ctx *xGinContext) initFlyTraceInfo() {
	ctx.flyTrace = xcontext.FlyTrace(GetTraceID(ctx.raw))
}

func (ctx *xGinContext) initBizTraceInfo() {
	ctx.bizTrace = xcontext.BizTrace(GetBizTraceID(ctx.raw))
}

func (ctx *xGinContext) initLang() {
	InitLang(ctx.raw)
	ctx.lang = xcontext.LangInfo{Code: GetLang(ctx.raw)}
}


func InitTraceID(ctx *gin.Context) {
	traceIDVal := ctx.GetHeader(CtxKey_TraceID)
	// 此处做兼容处理
	if traceIDVal == "" {
		traceIDVal = ctx.GetHeader(CtxKey_OldTraceID)
	}
	if traceIDVal == "" {
		traceIDVal = xstring.GetUUID()
	}
	ctx.Set(CtxKey_TraceID, traceIDVal)
}

func GetTraceID(ctx *gin.Context) string {
	if v, exists := ctx.Get(CtxKey_TraceID); exists {
		return v.(string)
	}

	return ""
}

func InitBizTraceID(ctx *gin.Context) {
	traceIDVal := ctx.GetHeader(CtxKey_ChainID)
	if traceIDVal == "" {
		traceIDVal = xstring.GetUUID()
	}
	ctx.Set(CtxKey_ChainID, traceIDVal)
}

func GetBizTraceID(ctx *gin.Context) string {
	if v, exists := ctx.Get(CtxKey_ChainID); exists {
		return v.(string)
	}

	return ""
}

// url>cookie> header > default
func InitLang(ctx *gin.Context) {
	var err error
	lang:= ctx.Query(CtxKey_QueryLang)
	if lang == ""{
		lang,err= ctx.Cookie(CtxKey_CookieLang)
		if err != nil && err != http.ErrNoCookie {
			xlog.Error(err.Error())
		}
	}
	if lang == ""{
		lang = ctx.GetHeader(CtxKey_HeaderLang)
	}
	if lang == ""{
		lang = i18n.GetDefaultLang()
	}
	if lang == ""{
		lang = xcontext.DEFAULT_LANG
	}
	ctx.Set(CtxKey_Lang,lang)
}

func GetLang(ctx *gin.Context) string {
	if v, exists := ctx.Get(CtxKey_Lang); exists {
		return v.(string)
	}
	return xcontext.DEFAULT_LANG
}


// ===== flyContext start ======
func (ctx *xGinContext) GetRawContext() interface{} {
	return ctx.raw
}

func (ctx *xGinContext) GetFlyTrace() xcontext.FlyTrace {
	return ctx.flyTrace
}

func (ctx *xGinContext) GetLang() xcontext.LangInfo {
	return ctx.lang
}

func (ctx *xGinContext) GetMeshTrace() xcontext.MeshTrace {
	return ctx.meshTrace
}

func (ctx *xGinContext) GetProtocol() xcontext.ProtocolType {
	return ctx.protocol
}

func (ctx *xGinContext) GetUserInfo() xcontext.UserInfo {
	return ctx.user
}

func (ctx *xGinContext) GetBizTrace() xcontext.BizTrace {
	return ctx.bizTrace
}

func (ctx *xGinContext) Get(key string) (value interface{}, exist bool) {
	return ctx.raw.Get(key)
}

func (ctx *xGinContext) Set(key string, value interface{}) {
	ctx.raw.Set(key, value)
}

func (ctx *xGinContext) Deadline() (deadline time.Time, ok bool) {
	deadline, ok = ctx.raw.Deadline()
	return
}

func (ctx *xGinContext) Done() <-chan struct{} {
	return ctx.raw.Done()
}

func (ctx *xGinContext) Err() error {
	return ctx.raw.Err()
}

func (ctx *xGinContext) Value(key interface{}) interface{} {
	return ctx.raw.Value(key)
}

// === flyContext end