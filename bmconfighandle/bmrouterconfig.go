package bmconfig

type BMRouterConfig struct {
	Host string
	Port string
}

func (br *BMRouterConfig) GenerateConfig() {
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	configPath := "resource/routerconfig.json"
	profileItems := BMGetConfigMap(configPath)

	br.Host = profileItems["Host"].(string)
	br.Port = profileItems["Port"].(string)
}