package conf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var g_config *goconfig.ConfigFile
func init() {
	config, err := goconfig.LoadConfigFile("conf/conf.ini")
	if err != nil {
		fmt.Println("load config file fail:", err)
	}
	g_config = config
}

func GetConfigSection(section string) map[string]string {
	secontion, err := g_config.GetSection(section)
	if err != nil {
		return nil
	}
	return secontion
}

func GetConfigValue (section, key string) string{
	value, err := g_config.GetValue(section, key)
	if err != nil {
		fmt.Println("get config value fail:", err)
		return ""
	}
	return value
}