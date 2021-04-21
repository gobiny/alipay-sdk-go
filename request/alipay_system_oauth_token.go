package request

const AlipaySystemOauthTokenMethod = "alipay.system.oauth.token"

type AlipaySystemOauthTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
