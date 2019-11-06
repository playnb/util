package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var defaultCfg *Config
var allCfgs []*Config

func init() {
	defaultCfg = &Config{Name: "default", cfg: viper.GetViper()}
}

func Default() *Config {
	return defaultCfg
}

func Init() {
	var err error
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(fmt.Errorf("Fatal error bind flag: %s \n", err))
	}

	for _, c := range allCfgs {
		c.cfg.SetConfigFile(viper.GetString(c.Name + ".file"))
		c.cfg.AddConfigPath(viper.GetString(c.Name + ".path"))
		err = c.cfg.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		c.cfg.WatchConfig()
		c.cfg.OnConfigChange(func(e fsnotify.Event) {
			c.onLoad()
		})
		c.onLoad()
	}
}

func New(name string) *Config {
	flag.String(name+".file", "config.toml", "读取的配置文件")
	flag.String(name+".path", ".", "读取的配置路径")
	cfg := &Config{Name: name, cfg: viper.New()}
	allCfgs = append(allCfgs, cfg)
	return cfg
}
