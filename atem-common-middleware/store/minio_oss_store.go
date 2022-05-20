package store

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioOssClient struct {
	option   *StoreConfig
	minioCli *minio.Client
}

func NewMinioOssClient(config *StoreConfig) *MinioOssClient {
	cli := &MinioOssClient{option: config}
	return cli
}

// Client : 创建oss client对象
func (c *MinioOssClient) GetClient() *minio.Client {
	if c.minioCli != nil {
		return c.minioCli
	}
	ossCli, err := minio.New(c.option.OSSEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.option.OSSAccesskey, c.option.OSSAccessSecret, ""),
		Secure: c.option.UseSLL,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	c.minioCli = ossCli
	return ossCli
}

func (c *MinioOssClient) AutoCreateBucket(bucketName string) error {
	ctx := context.Background()
	err := c.GetClient().MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := c.minioCli.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			return errors.New("We already own " + bucketName + "\n")
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
		return nil
	}
}

func (c *MinioOssClient) UploadObject(
	bucketName, objectName, contentType string,
	file *multipart.FileHeader) (info minio.UploadInfo, err error) {

	ctx := context.Background()
	src, err := file.Open()
	if err != nil {
		return
	}
	defer src.Close()
	if contentType == "" {
		contentType = "image/jpeg"
	}
	mcil := c.GetClient()
	// 使用PutObject上传一个zip文件
	return mcil.PutObject(ctx, bucketName, objectName, src, -1, minio.PutObjectOptions{ContentType: contentType})
}

func (c *MinioOssClient) FPutObject(
	bucketName, objectName, contentType,
	filePath string) (info minio.UploadInfo, err error) {
	ctx := context.Background()
	cli := c.GetClient()
	if cli != nil {
		if contentType == "" {
			contentType = "image/jpeg"
		}
		info, err := cli.FPutObject(ctx, bucketName, objectName, filePath,
			minio.PutObjectOptions{ContentType: contentType})
		if err == nil {
			return info, err
		}
	}
	return
}
