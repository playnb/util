package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Name      string
	cfg       *viper.Viper
	loadFuncs []func()
}

func (c *Config) AddLoadFunc(f func()) {
	c.loadFuncs = append(c.loadFuncs, f)
}
func (c *Config) Viper() *viper.Viper {
	return c.cfg
}
func (c *Config) onLoad() {
	for _, f := range c.loadFuncs {
		f()
	}
}

func (c *Config) Get(key string) interface{} {
	return c.cfg.Get(key)
}
func (c *Config) GetBool(key string) bool {
	return c.cfg.GetBool(key)
}
func (c *Config) GetFloat64(key string) float64 {
	return c.cfg.GetFloat64(key)
}
func (c *Config) GetInt(key string) int {
	return c.cfg.GetInt(key)
}
func (c *Config) GetIntSlice(key string) []int {
	return c.cfg.GetIntSlice(key)
}
func (c *Config) GetString(key string) string {
	return c.cfg.GetString(key)
}
func (c *Config) GetStringMap(key string) map[string]interface{} {
	return c.cfg.GetStringMap(key)
}
func (c *Config) GetStringMapString(key string) map[string]string {
	return c.cfg.GetStringMapString(key)
}
func (c *Config) GetStringSlice(key string) []string {
	return c.cfg.GetStringSlice(key)
}
func (c *Config) GetTime(key string) time.Time {
	return c.cfg.GetTime(key)
}
func (c *Config) GetDuration(key string) time.Duration {
	return c.cfg.GetDuration(key)
}
func (c *Config) IsSet(key string) bool {
	return c.cfg.IsSet(key)
}
func (c *Config) AllSettings() map[string]interface{} {
	return c.cfg.AllSettings()
}
