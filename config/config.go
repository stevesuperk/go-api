// 根据文件进行初始化配置，例如设定日志的配置，并监控文件的变化

package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 读取配置
type Config struct {
	Name string
}

var log *logrus.Logger

// Init 初始化配置，默认读取config.local.yaml
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	//c.initLog()

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (cfg *Config) initConfig() error {
	if cfg.Name != "" {
		viper.SetConfigFile(cfg.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("config/conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("local.yaml")
	}
	viper.SetConfigType("yaml")  // 设置配置文件格式为YAML
	viper.AutomaticEnv()         // 读取匹配的环境变量
	viper.SetEnvPrefix("go-api") // 读取环境变量的前缀为 snake
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return errors.WithStack(err)
	}

	return nil
}

// 监控配置文件变化并热加载程序
func (cfg *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Info("Config file changed: %s, \n %v", e.Name,viper.GetString("db.username"))
	})
}
