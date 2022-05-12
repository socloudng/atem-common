package captcha

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

func NewCaptchaRedisStore(rc *redis.Client, l *zap.Logger) *captchaRedisStore {
	return &captchaRedisStore{
		expiration: time.Second * 180,
		preKey:     "CAPTCHA_",
		logger:     l,
		redisCli:   rc,
	}
}

type captchaRedisStore struct {
	expiration time.Duration
	preKey     string
	context    context.Context
	redisCli   *redis.Client
	logger     *zap.Logger
}

// v8下使用redis
func (c *captchaRedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	c.context = ctx
	return c
}

func (c *captchaRedisStore) Set(id string, value string) error {
	err := c.redisCli.Set(c.context, c.preKey+id, value, c.expiration).Err()
	if err != nil {
		c.logger.Error("RedisStoreSetError!", zap.Error(err))
	}
	return err
}

func (c *captchaRedisStore) Get(key string, clear bool) string {
	val, err := c.redisCli.Get(c.context, key).Result()
	if err != nil {
		c.logger.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := c.redisCli.Del(c.context, key).Err()
		if err != nil {
			c.logger.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (c *captchaRedisStore) Verify(id, answer string, clear bool) bool {
	key := c.preKey + id
	v := c.Get(key, clear)
	return v == answer
}
