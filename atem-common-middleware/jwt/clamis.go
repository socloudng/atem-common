package jwt

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func (j *JWTService) GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	claims, err := j.ParseToken(token)
	if err != nil {
		j.Logger.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func (j *JWTService) GetUserID(c *gin.Context) uint64 {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.ID
	}

}

// 从Gin的Context中获取从jwt解析出来的用户UUID
func (j *JWTService) GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.UUID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func (j *JWTService) GetUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return ""
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.AuthorityId
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id
func (j *JWTService) GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := j.GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
