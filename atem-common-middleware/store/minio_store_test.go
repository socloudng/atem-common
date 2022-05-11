package store

import (
	"fmt"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	cfg := &StoreConfig{
		OSSEndpoint:     "192.168.1.156:9000",
		OSSAccesskey:    "visual",
		OSSAccessSecret: "visual17701305321",
		OSSBucket:       "default",
		UseSLL:          false,
	}
	cli := NewMinioOssClient(cfg)
	o := cli.AutoCreateBucket("haha")
	fmt.Println("=======" + o.Error() + "=======")
}
