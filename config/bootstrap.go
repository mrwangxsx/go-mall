package config

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"os"
	"time"
)

// **嵌入文件只能在写embed指令的Go文件的同级目录或者子目录中
//
//go:embed *.yaml
var configs embed.FS

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	vp := viper.New()
	// 根据环境变量ENV决定要读取的应用启动配置
	configFileStream, err := configs.ReadFile("application." + env + ".yaml")
	if err != nil {
		panic(err)
	}
	// 设置配置文件类型为 YAML
	vp.SetConfigType("yaml")
	// 将配置文件内容读取到 Viper 中
	err = vp.ReadConfig(bytes.NewBuffer(configFileStream))
	if err != nil {
		panic(err)
	}
	// 将配置文件中的 "app" 部分反序列化到 App 结构体
	vp.UnmarshalKey("app", &App)
	vp.UnmarshalKey("database", &Database)
	Database.MaxLifeTime *= time.Second
}
