package utils

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Log    *logrus.Logger
	Conifg *viper.Viper
)

func init() {
	Log = InitLogger()
	Conifg = InitConfig()
}

func InitConfig() *viper.Viper {
	// 设置 Viper 配置文件的名称（不带扩展名）
	viper.SetConfigName("config")
	// 设置 Viper 配置文件的类型
	viper.SetConfigType("yaml")
	// 添加 Viper 配置文件的搜索路径，例如当前目录
	viper.AddConfigPath(".")

	// 尝试读取配置文件
	err := viper.ReadInConfig()

	// 如果配置文件不存在，则创建一个新的配置文件
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			Log.Debug("找不到配置文件，将创建新的配置文件。")

			// 创建新的配置文件
			err := createConfigFile()
			if err != nil {
				Log.Info(fmt.Sprintf("无法创建配置文件: %s\n", err))
				return nil
			}

			// 重新尝试读取配置文件
			err = viper.ReadInConfig()
			if err != nil {
				Log.Info(fmt.Sprintf("无法读取配置文件: %s\n", err))
				return nil
			}
		} else {
			Log.Info(fmt.Sprintf("无法读取配置文件: %s\n", err))
			return nil
		}
	}

	// 检查必要的字段是否已填写
	checkRequiredFields("WeatherKey", "CnameData", "CnameImage", "AppID", "AppSecret")

	return viper.GetViper()
}

// createConfigFile 创建新的配置文件并写入默认值
func createConfigFile() error {
	// 设置默认配置项
	viper.SetDefault("WeatherKey", "")
	viper.SetDefault("CnameData", "")
	viper.SetDefault("CnameImage", "")
	viper.SetDefault("AppID", "")
	viper.SetDefault("AppSecret", "")

	// 保存配置到文件
	err := viper.WriteConfigAs("config.yaml")
	if err != nil {
		return err
	}

	Log.Debug("配置文件已创建。请编辑 config.yaml 文件并填写相应的字段值。")
	os.Exit(0)

	return nil
}

// checkRequiredFields 检查必要的字段是否已填写
func checkRequiredFields(fields ...string) {
	missingFields := []string{}

	for _, field := range fields {
		if !viper.IsSet(field) || viper.GetString(field) == "" {
			missingFields = append(missingFields, field)
		}
	}

	if len(missingFields) > 0 {
		Log.Info(fmt.Sprintf("配置文件缺少必要字段: %v，请填写这些字段后重试。\n", missingFields))
		os.Exit(0)
	}
}
