package bmxmpp

import (
	"blackmirror/bmconfighandle"
	"os"
	"sync"
)

type BmXmppConfig struct {
	Host          string
	Port          string
	HostName      string
	LoginUser     string
	LoginUserPwd  string
	ReportUser    string
	ReportUserPwd string
	ListenUser    string
	ListenUserPwd string
}

var e error
var onceConfig sync.Once
var config *BmXmppConfig

func GetConfigInstance() (*BmXmppConfig, error) {
	onceConfig.Do(func() {
		configPath := os.Getenv("BM_XMPP_CONF_HOME")
		profileItems := bmconfig.BMGetConfigMap(configPath)
		config = &BmXmppConfig{
			Host:profileItems["Host"].(string),
			Port:profileItems["Port"].(string),
			HostName:profileItems["HostName"].(string),
			LoginUser:profileItems["LoginUser"].(string),
			LoginUserPwd:profileItems["LoginUserPwd"].(string),
			ReportUser:profileItems["ReportUser"].(string),
			ReportUserPwd:profileItems["ReportUserPwd"].(string),
			ListenUser:profileItems["ListenUser"].(string),
			ListenUserPwd:profileItems["ListenUserPwd"].(string),

		}
		e = nil
	})
	return config, e
}
