package i18n

const (
	//i18n配置Key常量
	OptionKeyDefaultLang  string = "i18n.default.lang"
	OptionKeyDataProvider string = "i18n.data.provider"
	OptionKeyCacheMode    string = "i18n.cache.mode"
	OptionKeyCacheCron    string = "i18n.cache.cron"

	//i18n配置值常量
	OptionValueDefaultLangDefault string = "zh-CN"
	OptionValueDataProviderNone   string = `{"mode":"none"}`
	OptionValueCacheModeRefresh   string = "refresh"
	OptionValueCacheModeClear     string = "clear"
	OptionValueCacheCronDefault   string = "12 34 * * * ?"
)



// GetDefaultLang 获取配置的默认语言。
func GetDefaultLang() string {
	/*if options == nil {
		return ""
	}

	return options.DefaultLang*/

	return OptionValueDefaultLangDefault
}
