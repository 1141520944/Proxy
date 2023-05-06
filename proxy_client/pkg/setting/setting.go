package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局的变量，用来保存程序的所有的配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name       string `mapstructure:"name"`
	Mode       string `mapstructure:"mode"`
	Version    string `mapstructure:"version"`
	Port       int    `mapstructure:"port"`
	StartTime  string `mapstructure:"start_time"`
	MachineID  int64  `mapstructure:"machine_id"`
	RemoteIp   string `mapstructure:"remote_ip"`
	RemotePort string `mapstructure:"remote_port"`
	*LogConfig `mapstructure:"log"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// Init viper读取文件
func Init() (err error) {
	//指定文件名称
	viper.SetConfigName("config")
	//指定文件类型
	viper.SetConfigType("yaml")
	//指定文件路径
	viper.AddConfigPath("./config")
	//处理配置文件
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}
	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//热加载-配置文件实时的监控配置
	viper.WatchConfig()
	//钩子函数-当config改变时：
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})
	return
}
