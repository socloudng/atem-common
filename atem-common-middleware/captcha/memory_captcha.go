package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
)

func NewCaptchaMemStore() *captchaMemStore {
	return &captchaMemStore{
		store: base64Captcha.DefaultMemStore,
	}
}

type captchaMemStore struct {
	store base64Captcha.Store
}

func (c *captchaMemStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	return c
}

func (c *captchaMemStore) Set(id string, value string) error {
	return c.store.Set(id, value)
}

func (c *captchaMemStore) Get(key string, clear bool) string {
	return c.store.Get(key, clear)
}

func (c *captchaMemStore) Verify(id, answer string, clear bool) bool {
	return c.store.Verify(id, answer, clear)
}
