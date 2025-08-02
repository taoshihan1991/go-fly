package models

var CustomConfigs []Config

type Config struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	ConfName  string `json:"conf_name"`
	ConfKey   string `json:"conf_key"`
	ConfValue string `json:"conf_value"`
	UserId    string `json:"user_id"`
}

func UpdateConfig(userid interface{}, key string, value string) {
	config := FindConfigByUserId(userid, key)
	if config.ID != 0 {
		config.ConfValue = value
		DB.Model(&Config{}).Where("user_id = ? and conf_key = ?", userid, key).Update(config)
	} else {
		newConfig := &Config{
			ID:        0,
			ConfName:  "",
			ConfKey:   key,
			ConfValue: value,
			UserId:    userid.(string),
		}
		DB.Create(newConfig)
	}

}
func FindConfigs() []Config {
	var config []Config
	DB.Find(&config)
	return config
}
func FindConfigsByUserId(userid interface{}) []Config {
	var config []Config
	DB.Where("user_id = ?", userid).Find(&config)
	return config
}

func FindConfig(key string) string {
	for _, config := range CustomConfigs {
		if key == config.ConfKey {
			return config.ConfValue
		}
	}
	return ""
}
func FindConfigByUserId(userId interface{}, key string) Config {
	var config Config
	DB.Where("user_id = ? and conf_key = ?", userId, key).Find(&config)
	return config
}
