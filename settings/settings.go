package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量
var Conf = new(AppConfig)

type AppConfig struct {
	Mode       string `mapstructure:"mode"`
	*LogConfig `mapstructure:"log"`
	*K8sInfo   `mapstructure:"k8sinfo"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type K8sInfo struct {
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Token string `mapstructure:"token"`
}

func Init() (err error) {

	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件
	err = viper.ReadInConfig()                // 读取配置信息
	if err != nil {                           // 读取配置信息失败
		fmt.Printf("viper failed,%v\n", err)
		return
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了.....")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}

	})
	return err
}
