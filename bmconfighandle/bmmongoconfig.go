package bmconfig

type BMMongoConfig struct {
	Host string
	Port string
	Database string
}

func (mc *BMMongoConfig) GenerateConfig() {
	//TODO: 配置文件路径 待 用脚本指定dev路径和deploy路径
	configPath := "resource/mongoconfig.json"
	profileItems := BMGetConfigMap(configPath)

	mc.Host = profileItems["Host"].(string)
	mc.Port = profileItems["Port"].(string)
	mc.Database = profileItems["Database"].(string)

}
