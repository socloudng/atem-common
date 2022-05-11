package jwt

import (
	"atem/atem-common/atem-common-base/base_model"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type JwtBlackList struct {
	base_model.BASE_MODEL[uint64]
	Jwt string `gorm:"type:text;comment:jwt"`
}

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint64
	Username    string
	NickName    string
	AuthorityId string
}
