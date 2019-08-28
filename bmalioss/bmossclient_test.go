package bmalioss

import (
	"fmt"
	"blackmirror/bmerror"
	"github.com/hashicorp/go-uuid"
	"io/ioutil"
	"os"
	"testing"
)

func TestPushObject(t *testing.T) {
	os.Setenv("BM_OSS_CONF_HOME", "../resource/ossconfig.json")

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
