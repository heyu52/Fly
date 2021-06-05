package xgin

import (
	"github.com/gin-gonic/gin"
	"github.com/heyu52/Fly/contract"
	"github.com/heyu52/Fly/xcontext"
	"github.com/heyu52/Fly/xhttp"
)

const (
	ctxKeyPref = xcontext.FLY_PREFIX

	// auth
	CtxKey_AuthTypeKey = ctxKeyPref + "AuthType"
	CtxKey_User        = xcontext.USER_KEY
	CtxKey_UserID      = xcontext.USERID_KEY
	CtxKey_UserName    = xcontext.USERNAME_KEY
	// servelog 堆栈
	CtxKey_Stack = ctxKeyPref + "Stack"
	// servelog 服务产生的真实错误信息
	CtxKey_RealError = ctxKeyPref + "RealError"
	CtxKey_TraceID   = xcontext.RASSE_TRACE_KEY
	CtxKey_ChainID   = xcontext.BIZ_TRACE_KEY
	// Deprecated: Use CtxKey_TraceID instead.
	CtxKey_OldTraceID = "rasse-trace-id"
	CtxKey_QueryLang   = "lang"
	// 参考规范：https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Language
	CtxKey_HeaderLang = xhttp.LangHeaderKey
	CtxKey_CookieLang = xcontext.LANG_KEY
	CtxKey_Lang = xcontext.LANG_KEY
)

func GetCurrentUser(ctx *gin.Context) *contract.Staff {
 if user,ok:=ctx.Keys[CtxKey_User];ok && user !=nil{
 	return user.(*contract.Staff)
 }else {
	var userId=ctx.GetHeader(CtxKey_UserID)
	var userName=ctx.GetString(CtxKey_UserName)
	if user!="" && userName!=""{
		return &contract.Staff{
			StaffID: userId,
			StaffName: userName,
		}
	}
 }
 return nil
}
