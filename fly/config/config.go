package config

import (
	"fmt"
	"log"

	"github.com/riba2534/let_the_bullets_fly_ai/fly/utils"
	"github.com/spf13/viper"
)

type Config struct {
	OPENAI_URL     string `mapstructure:"OPENAI_URL"`
	OPENAI_API_KEY string `mapstructure:"OPENAI_API_KEY"`
	QDRANT_URL     string `mapstructure:"QDRANT_URL"`
	QDRANT_PORT    string `mapstructure:"QDRANT_PORT"`
	QDRANT_API_KEY string `mapstructure:"QDRANT_API_KEY"`
}

var C Config

func Init() {
	// 设置配置文件路径和文件名
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	err := viper.ReadInConfig() // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败：%w", err))
	}

	err = viper.Unmarshal(&C) // 将配置文件绑定到结构体上
	if err != nil {
		panic(fmt.Errorf("绑定配置文件到结构体失败：%w", err))
	}

	// 打印读取到的配置信息
	log.Printf("读取配置成功:\n%s", utils.MarshalAnyToString(C))
}
