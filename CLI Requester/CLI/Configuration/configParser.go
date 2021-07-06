package Configuration

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ProjectID string
	AccessKey string
	SecretKey string
}

func LoadConfig(configPath string) (Config, bool) {
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	var con Config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&con)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	if len(con.ProjectID) == 0 {
		log.Printf("Config does not have ProjectID field")
		return Config{}, true
	}
	if len(con.AccessKey) == 0 {
		log.Printf("Config does not have AccessKey field")
		return Config{}, true
	}
	if len(con.SecretKey) == 0 {
		log.Printf("Config does not have SecretKey field")
		return Config{}, true
	}

	//log.Printf("ProjectID: %s", con.ProjectID)
	//log.Printf("AccessKey: %s", con.AccessKey)
	//log.Printf("SecretKey: %s", con.SecretKey)
	return con, false
}
