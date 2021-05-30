package contract

const (
	CODE_OK     = 0
	CODE_BIZ    = 1
	CODE_SERVER = -1

	// 租户相关错误
	CODE_TENANCY_ERROR = 1000
	// 权限相关错误
	CODE_AUTH_ERROR = 10000
	// 参数相关错误
	CODE_PARAMS_ERROR = 20000
	// 业务相关错误
	CODE_BIZ_ERROR = 30000
	// 资源未找到
	CODE_NOTFOUND_ERROR = 40000
)

// ApiResult 接口返回体
type ApiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Stack   string      `json:"stack,omitempty"`
}

func NewApiResult(code int, msg string, data interface{}) *ApiResult {
	return &ApiResult{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

func OK(data interface{}) *ApiResult {
	return NewApiResult(0, "", data)
}

func Error(code int, msg string) *ApiResult {
	return NewApiResult(code, msg, nil)
}