package store

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

type CephClient struct {
	options  *CephConfig
	cephConn *s3.S3
}

func NewCethClient(config *CephConfig) *CephClient {
	return &CephClient{options: config}
}

// GetCephConnection : 获取ceph连接
func (c *CephClient) GetCephConnection() *s3.S3 {
	if c.cephConn != nil {
		return c.cephConn
	}
	// 1. 初始化ceph的一些信息

	auth := aws.Auth{
		AccessKey: c.options.CephAccessKey,
		SecretKey: c.options.CephSecretKey,
	}

	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          c.options.CephGWEndpoint,
		S3Endpoint:           c.options.CephGWEndpoint,
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}

	// 2. 创建S3类型的连接
	return s3.New(auth, curRegion)
}

// GetCephBucket : 获取指定的bucket对象
func (c *CephClient) GetCephBucket(bucket string) *s3.Bucket {
	conn := c.GetCephConnection()
	return conn.Bucket(bucket)
}

// PutObject : 上传文件到ceph集群
func (c *CephClient) PutObject(bucket string, path string, data []byte) error {
	return c.GetCephBucket(bucket).Put(path, data, "octet-stream", s3.PublicRead)
}
