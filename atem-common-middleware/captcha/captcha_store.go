package captcha

import (
	"context"

	"github.com/mojocn/base64Captcha"
)

type CaptchaStore interface {
	base64Captcha.Store
	UseWithCtx(ctx context.Context) base64Captcha.Store
}
