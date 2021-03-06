package alipaysdk

const (
	ServerUrlProduction  = "https://openapi.alipay.com/gateway.do"    // 正式环境请求地址
	ServerUrlDevelopment = "https://openapi.alipaydev.com/gateway.do" // 沙箱环境请求地址

	PublicAppAuthUrlProduction  = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"    // 正式环境授权登录地址
	PublicAppAuthUrlDevelopment = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm" // 沙箱环境授权登录地址

	AppToAppAuthUrlProduction  = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"    // 正式环境第三方授权登录地址
	AppToAppAuthUrlDevelopment = "https://openauth.alipaydev.com/oauth2/appToAppAuth.htm" // 沙箱环境第三方授权登录地址
)
