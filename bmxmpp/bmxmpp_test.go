package bmxmpp

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"os"
	"testing"
)

func TestKafkaProducer(t *testing.T) {

	os.Setenv("BM_XMPP_CONF_HOME", "../resource/xmppconfig.json")

	bxc, err := GetConfigInstance()
	bmerror.PanicError(err)
	err = bxc.Forward("test@max.logic", "test func")
	bmerror.PanicError(err)

}
