package bmalioss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
	"os"
)

func PushOneObject(buket string, name string, path string) error {

	sts, err := QuerySTSToken()
	if err != nil {
		return err
	}

	cliopt := func (client *oss.Client) {
		client.Config.SecurityToken = sts.GetSecurityToken()
	}

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", sts.GetAccessKeyId(), sts.GetAccessKeySecret(), cliopt)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(buket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile(name, path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	return err
}
