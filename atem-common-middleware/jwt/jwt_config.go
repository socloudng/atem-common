package jwt

import "errors"

type JwtConfig struct {
	SigningKey    string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`          // jwt签名
	ExpiresTime   int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"`       // 过期时间
	BufferTime    int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`          // 缓冲时间
	Issuer        string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                        // 签发者
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // 多点登录拦截
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)
