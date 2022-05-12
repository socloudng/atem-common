package configs

type CaptchaConfig struct {
	StoreType string `mapstructure:"store" yaml:"store"`
	KeyLong   int    `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`       // 验证码长度
	ImgWidth  int    `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`    // 验证码宽度
	ImgHeight int    `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"` // 验证码高度
}
