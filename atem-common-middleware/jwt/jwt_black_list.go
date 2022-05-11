package jwt

import (
	"context"
	"time"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

var blackCache local_cache.Cache

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JWTService) GetRedisJWT(userName string) (string, error) {
	return jwtService.Redis.Get(context.Background(), userName).Result()
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JWTService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(jwtService.option.ExpiresTime) * time.Second
	err = jwtService.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JWTService) IsBlacklist(jwt string) bool {
	_, ok := jwtService.getBlackList().Get(jwt)
	return ok
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JWTService) JsonInBlacklist(jwtList JwtBlackList) (err error) {
	err = jwtService.Orm.Create(&jwtList).Error
	if err != nil {
		return
	}
	jwtService.getBlackList().SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (jwtService *JWTService) getBlackList() *local_cache.Cache {
	if (local_cache.Cache{}) == blackCache {
		jwtService.loadAllBlackList()
	}
	return &blackCache
}

func (jwtService *JWTService) loadAllBlackList() {
	if jwtService.option == nil {
		jwtService.Logger.Fatal("please set jwt-config first")
	}
	blackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(jwtService.option.ExpiresTime)),
	)
	var data []string
	err := jwtService.Orm.Model(&JwtBlackList{}).Select("jwt").Find(&data).Error
	if err != nil {
		jwtService.Logger.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		blackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
