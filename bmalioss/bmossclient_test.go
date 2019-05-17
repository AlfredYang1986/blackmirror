package bmalioss

import (
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"os"
	"testing"
)

func TestPushObject(t *testing.T) {
	os.Setenv("BM_OSS_CONF_HOME", "../resource/ossconfig.json")
	os.Setenv("BM_OSS_TEMP_DIR", "/home/jeorch/work/test/temp")

	bucketName := "pharbers-resources"
	tempUUID,_ := uuid.GenerateUUID()
	fmt.Println(tempUUID)
	objectKey := tempUUID

	localDir := "/home/jeorch/work/test/temp/test.jpeg"
	f, err := os.Open(localDir)
	bmerror.PanicError(err)
	defer f.Close()

	objectValue, err := ioutil.ReadAll(f)
	bmerror.PanicError(err)

	err = PutObject(bucketName, objectKey, objectValue)
	bmerror.PanicError(err)
}
