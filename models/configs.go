package models

var CustomConfigs []Config
type Config struct{
	ID     uint `gorm:"primary_key" json:"id"`
	ConfName string `json:"conf_name"`
	ConfKey string `json:"conf_key"`
	ConfValue string `json:"conf_value"`
}

func FindConfigs()[]Config{
	var config []Config
	DB.Find(&config)
	return config
}
func InitConfig(){
	CustomConfigs=FindConfigs()
}
func FindConfig(key string)string{
	for _,config:=range CustomConfigs{
		if key==config.ConfKey{
			return config.ConfValue
		}
	}
	return ""
}