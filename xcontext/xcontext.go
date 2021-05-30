package xcontext

import (
	"context"
	"github.com/spf13/viper"
	"strings"
)

// rasse 上下文，与协议无关。
type XContext interface {
	context.Context
	// 获取原生的context
	GetRawContext() interface{}
	// 获取当前协议
	GetProtocol() ProtocolType
	// 获取全链路跟踪信息
	GetMeshTrace() MeshTrace
	// 获取请求跟踪信息
	GetRasseTrace() RasseTrace
	//  获取用户信息
	GetUserInfo() UserInfo
	// 获取语言信息
	GetLang() LangInfo
	// 获取业务链路跟踪信息
	GetBizTrace() BizTrace
	//  从上下文中获取内容
	Get(key string) (value interface{}, exist bool)
	// 给上下文赋值
	Set(key string, value interface{})
}

// 协议
type ProtocolType int32

type LangInfo struct {
	Code string `json:"code"`
}

const (
	_ ProtocolType = iota
	Gin
	Trpc
	Test
	TDMQ
)

func (p ProtocolType) String() string {
	switch p {
	case Gin:
		return "GIN"
	case Trpc:
		return "TRPC"
	case Test:
		return "TEST"
	case TDMQ:
		return "TDMQ"
	default:
		return "UNKNOWN"
	}
}

type BizTrace string

type RasseTrace string

type MeshTrace map[string]string

type UserInfo struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

func ToMap(ctx XContext) map[string]string {

	var userInfo = ctx.GetUserInfo()
	var fields = make(map[string]string)
	fields[USERID_KEY] = userInfo.UserId
	fields[USERNAME_KEY] = userInfo.UserName
	fields[RASSE_TRACE_KEY] = string(ctx.GetRasseTrace())
	fields[BIZ_TRACE_KEY] = string(ctx.GetBizTrace())
	fields[LANG_KEY] = ctx.GetLang().Code
	apm := ctx.GetMeshTrace()
	meshKeys := GetMeshKeys()
	for _, k := range meshKeys {
		if apm != nil {
			fields[k] = apm[k]
		} else {
			fields[k] = ""
		}
	}
	return fields
}

func GetMeshKeys() []string {
	var headerKeys []string
	tmp := viper.GetString(MESH_CONFIG_KEY)
	if tmp != "" {
		headerKeys = strings.Split(tmp, ",")
	} else {
		headerKeys = append(headerKeys, DEFAULT_MESH_TRACE_KEYS...)
	}
	return headerKeys
}
