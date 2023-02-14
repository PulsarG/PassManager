package confile

import (
	"fmt"

	"github.com/go-ini/ini"
)

func CfgHandler() *ini.File {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return nil
	}
	return cfg
}

func GetFromIni(section, key string) string {
	return CfgHandler().Section(section).Key(key).String()
}

func SaveToIni(section, key, val string) {
	cfg := CfgHandler()
	cfg.Section(section).Key(key).SetValue(val)
	err := cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Printf("Error saving config file: %s\n", err)
		return
	}
}
