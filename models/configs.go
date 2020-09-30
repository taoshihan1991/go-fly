package models

var CustomConfigs []Config

type Config struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	ConfName  string `json:"conf_name"`
	ConfKey   string `json:"conf_key"`
	ConfValue string `json:"conf_value"`
}

func UpdateConfig(key string, value string) {
	c := &Config{
		ConfValue: value,
	}
	DB.Model(c).Where("conf_key = ?", key).Update(c)
	InitConfig()
}
func FindConfigs() []Config {
	var config []Config
	DB.Find(&config)
	return config
}
func InitConfig() {
	CustomConfigs = FindConfigs()
}
func FindConfig(key string) string {
	for _, config := range CustomConfigs {
		if key == config.ConfKey {
			return config.ConfValue
		}
	}
	return ""
}
