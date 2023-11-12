package utils

import (
	"os"

	"github.com/spf13/viper"
)

var Config struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func init() {
	// 检查配置文件是否存在
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		// 配置文件不存在，创建默认配置文件
		createDefaultConfig()
	}

	// 初始化 Viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		Log.Errorf("failed to read config file: %s", err)
		os.Exit(0)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		Log.Errorf("failed to unmarshal config: %s", err)
		os.Exit(0)
	}
}

func createDefaultConfig() {
	// 创建默认配置文件
	configContent := "username: your_username \npassword: your_password"
	err := os.WriteFile("config.yaml", []byte(configContent), 0644)
	if err != nil {
		Log.Errorf("failed to create config file: %s", err)
		os.Exit(0)
	}
}
