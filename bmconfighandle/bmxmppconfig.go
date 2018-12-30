package bmconfig

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

func (xc *BmXmppConfig) GenerateConfig() {
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	configPath := "resource/xmppconfig.json"
	profileItems := BMGetConfigMap(configPath)

	xc.Host = profileItems["Host"].(string)
	xc.Port = profileItems["Port"].(string)
	xc.HostName = profileItems["HostName"].(string)
	xc.LoginUser = profileItems["LoginUser"].(string)
	xc.LoginUserPwd = profileItems["LoginUserPwd"].(string)
	xc.ReportUser = profileItems["ReportUser"].(string)
	xc.ReportUserPwd = profileItems["ReportUserPwd"].(string)
	xc.ListenUser = profileItems["ListenUser"].(string)
	xc.ListenUserPwd = profileItems["ListenUserPwd"].(string)

}
