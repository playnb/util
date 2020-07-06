package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func CreateViperConfig(name string) *ViperConfig {
	c := &ViperConfig{}
	c.prefix = ""
	v := viper.New()
	c.parent = v
	c.name = strings.TrimRight(name, ".toml")
	return c
}

func GenViperConfig(v Config, prefix string) Config {
	c := &ViperConfig{}
	c.prefix = prefix
	c.parent = v
	return c
}

type ViperConfig struct {
	name     string
	prefix   string
	parent   Config
	loadFunc []func(Config)
	loaded   bool
}

func (c *ViperConfig) onLoad() {
	for _, f := range c.loadFunc {
		f(c)
	}
}
func (c *ViperConfig) key(k string) string {
	if len(c.prefix) > 0 {
		return c.prefix + "." + k
	}
	return k
}

func (c *ViperConfig) AddLoadFunc(f func(Config)) {
	if c == c.parent {
		panic("nest config define")
		return
	}
	switch v := c.parent.(type) {
	case *viper.Viper:
		c.loadFunc = append(c.loadFunc, f)
	case *ViperConfig:
		v.AddLoadFunc(f)
	}

}
func (c *ViperConfig) Viper() *viper.Viper {
	if c == c.parent {
		panic("nest config define")
		return nil
	}
	switch v := c.parent.(type) {
	case *viper.Viper:
		return v
	case *ViperConfig:
		return v.Viper()
	}
	return nil
}
func (c *ViperConfig) Load() {
	if c.loaded {
		return
	}
	c.loaded = true
	v := c.Viper()
	v.SetConfigFile(c.name + ".toml")
	v.AddConfigPath("./")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		c.onLoad()
	})
	c.onLoad()

}
func (c *ViperConfig) Get(key string) interface{} {
	return c.parent.Get(c.key(key))
}
func (c *ViperConfig) GetBool(key string) bool {
	return c.parent.GetBool(c.key(key))
}
func (c *ViperConfig) GetFloat64(key string) float64 {
	return c.parent.GetFloat64(c.key(key))
}
func (c *ViperConfig) GetInt(key string) int {
	return c.parent.GetInt(c.key(key))
}
func (c *ViperConfig) GetIntSlice(key string) []int {
	return c.parent.GetIntSlice(c.key(key))
}
func (c *ViperConfig) GetString(key string) string {
	return c.parent.GetString(c.key(key))
}
func (c *ViperConfig) GetStringMap(key string) map[string]interface{} {
	return c.parent.GetStringMap(c.key(key))
}
func (c *ViperConfig) GetStringMapString(key string) map[string]string {
	return c.parent.GetStringMapString(c.key(key))
}
func (c *ViperConfig) GetStringSlice(key string) []string {
	return c.parent.GetStringSlice(c.key(key))
}
func (c *ViperConfig) GetTime(key string) time.Time {
	return c.parent.GetTime(c.key(key))
}
func (c *ViperConfig) GetDuration(key string) time.Duration {
	return c.parent.GetDuration(c.key(key))
}
func (c *ViperConfig) IsSet(key string) bool {
	return c.parent.IsSet(c.key(key))
}
