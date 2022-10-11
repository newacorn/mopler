package mopler

// Configuration 必要的配置
type Configuration struct {
	ClientId     string //您应用的AppKey。
	ClientSecret string //您应用的SecretKey
	//您通过鉴权的方式获取到的access_token， 如果设置了，那么就会跳过获取token的步骤，并且需要您自己实现过期后刷新token的方法
	//设置此参数后，TokenMode和AuthMethod可以在配置初始化的时候不进行传递
	AccessToken string
	//AuthMode     string //授权模式，verification_url（验证码） 或者 qrcode_url（二维码）
	TokenMode string //token验证方式
	//根据授权模式来引导用户授权
	AuthMethod func(codeResp *Code)
}

type ConfigurationOption func(*Configuration)

func NewConfiguration(clientId string, clientSecret string, options ...ConfigurationOption) Configuration {
	c := Configuration{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	for _, option := range options {
		option(&c)
	}
	return c
}

func WithConfigurationAccessToken(accessToken string) ConfigurationOption {
	return func(c *Configuration) {
		c.AccessToken = accessToken
	}
}

func WithConfigurationTokenMode(tokenMode string) ConfigurationOption {
	return func(c *Configuration) {
		c.TokenMode = tokenMode
	}
}

func WithConfigurationAuthMethod(authMethod func(codeResp *Code)) ConfigurationOption {
	return func(c *Configuration) {
		c.AuthMethod = authMethod
	}
}
