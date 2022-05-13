package store

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunOssClient struct {
	option *StoreConfig
	cli    *oss.Client
}

func NewOssClient(config *StoreConfig) *AliyunOssClient {
	return &AliyunOssClient{option: config}
}

// Client : 创建oss client对象
func (c *AliyunOssClient) GetClient() *oss.Client {
	if c.cli != nil {
		return c.cli
	}
	ossCli, err := oss.New(c.option.OSSEndpoint,
		c.option.OSSAccesskey, c.option.OSSAccessSecret)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	c.cli = ossCli
	return ossCli
}

// Bucket : 获取bucket存储空间
func (c *AliyunOssClient) GetBucket() *oss.Bucket {
	cli := c.GetClient()
	if cli != nil {
		bucket, err := cli.Bucket(c.option.OSSBucket)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return bucket
	}
	return nil
}

// DownloadURL : 临时授权下载url
func (c *AliyunOssClient) GetDownloadURL(objName string) string {
	signedURL, err := c.GetBucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedURL
}

// BuildLifecycleRule : 针对指定bucket设置生命周期规则
func (c *AliyunOssClient) GetBuildLifecycleRule(bucketName string) {
	// 表示前缀为test的对象(文件)距最后修改时间30天后过期。
	ruleTest1 := oss.BuildLifecycleRuleByDays("rule1", "test/", true, 30)
	rules := []oss.LifecycleRule{ruleTest1}

	c.GetClient().SetBucketLifecycle(bucketName, rules)
}

// GenFileMeta :  构造文件元信息
func (c *AliyunOssClient) GenFileMeta(metas map[string]string) []oss.Option {
	options := []oss.Option{}
	for k, v := range metas {
		options = append(options, oss.Meta(k, v))
	}
	return options
}
