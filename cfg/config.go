package cfg

import (
	"github.com/spf13/viper"
	"log"
	"wxlogin/models"
)

var Config models.Config

func Init() {
	viper.AddConfigPath(".")      // 设置配置文件路径为当前目录
	viper.SetConfigName("config") // 设置配置文件的名字为 config（不需要文件扩展名）
	viper.SetConfigType("yaml")   // 设置配置文件类型为 YAML

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err = viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

}
