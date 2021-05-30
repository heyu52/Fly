package xcontext

const (
	// 统一前缀
	RASSE_PREFIX = "x-Fly-"
	// 业务链路跟踪Key
	BIZ_TRACE_KEY = RASSE_PREFIX + "biz-id"
	// 用户信息
	USER_KEY = RASSE_PREFIX + "user"
	// 用户信息
	USERNAME_KEY = RASSE_PREFIX + "user-name"
	// 用户信息
	USERID_KEY = RASSE_PREFIX + "user-id"
	// 请求唯一ID
	RASSE_TRACE_KEY = RASSE_PREFIX + "trace-id"
	// 全链路跟踪配置
	MESH_CONFIG_KEY = "mesh.headers"
	// 多语言
	LANG_KEY = RASSE_PREFIX + "lang"
)

// 默认的Mesh 的APM KEY
var DEFAULT_MESH_TRACE_KEYS = []string{"x-request-id", "x-b3-traceid", "x-b3-spanid", "x-b3-parentspanid",
	"x-b3-sampled", "x-b3-flags", "x-ot-span-xcontext", "x-envoy-attempt-count", "x-envoy-external-address", "x-envoy-original-path"}

// 参考规范：https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Language
var DEFAULT_LANG = "zh-CN"
