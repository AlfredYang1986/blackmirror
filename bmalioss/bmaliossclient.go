package bmalioss

import (
	"bytes"
	"blackmirror/bmconfighandle"
	"blackmirror/bmerror"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"sync"
)



var e error
var onceConfig sync.Once
var ossClient *oss.Client

func getClientInstance() (*oss.Client, error) {
	onceConfig.Do(func() {
		configPath := os.Getenv("BM_OSS_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)
		client, err := oss.New(profileItems["endpoint"].(string), profileItems["accessKey"].(string), profileItems["accessKeySecret"].(string))
		if err == nil {
			ossClient = client
		} else {
			e = err
		}
	})
	return ossClient, e
}

func PutObject(bucketName string, objectKey string, objectValue []byte) error {
	oc, err := getClientInstance()
	bmerror.PanicError(err)
	bucket, err := oc.Bucket(bucketName)
	bmerror.PanicError(err)
	err = bucket.PutObject(objectKey, bytes.NewReader(objectValue))
	return err
}
