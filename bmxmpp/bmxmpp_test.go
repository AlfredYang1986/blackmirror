package bmxmpp

import (
	"github.com/alfredyang1986/blackmirror/bmerror"
	"os"
	"testing"
)

func TestBmXmppConfig_Forward(t *testing.T) {

	os.Setenv("BM_XMPP_CONF_HOME", "../resource/xmppconfig.json")

	bxc, err := GetConfigInstance()
	bmerror.PanicError(err)
	err = bxc.Forward("test@max.logic", "test func 123")
	bmerror.PanicError(err)

}

func TestBmXmppConfig_Forward2Group(t *testing.T) {
	os.Setenv("BM_XMPP_CONF_HOME", "../resource/xmppconfig.json")
	bxc, err := GetConfigInstance()
	bmerror.PanicError(err)
	err = bxc.Forward2Group("troom@conference.max.logic", "test-group func 1")
	bmerror.PanicError(err)
}
