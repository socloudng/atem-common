package jwt

import (
	"github/socloudng/atem-common/atem-common-base/base_service"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	signingKey []byte
	base_service.BaseService
	option *JwtConfig
}

func (j *JWTService) SetConfig(config *JwtConfig) {
	j.option = config
}

func (j *JWTService) CreateClaims(baseClaims BaseClaims) CustomClaims {
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: j.option.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                 // 签名生效时间
			ExpiresAt: time.Now().Unix() + j.option.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    j.option.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWTService) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.getSigningKey())
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWTService) CreateTokenByOldToken(oldToken string, claims CustomClaims) (string, error) {
	v, err, _ := j.Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWTService) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.getSigningKey(), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func (j *JWTService) getSigningKey() []byte {
	if j.option == nil {
		j.Logger.Fatal("please set jwt-config first")
	}
	if len(j.signingKey) < 1 {
		j.signingKey = []byte(j.option.SigningKey)
	}
	return j.signingKey
}
