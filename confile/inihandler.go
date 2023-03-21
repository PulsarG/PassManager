// HANDLER CONFIG FILE
// READ, WHRITE, TAKE

package confile

import (
	"fmt"

	"github.com/go-ini/ini"
)

func cfgHandler() *ini.File {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return nil
	}
	return cfg
}

func GetFromIni(section, key string) string {
	return cfgHandler().Section(section).Key(key).String()
}

func SaveToIni(section, key, val string) {
	cfg := cfgHandler()
	cfg.Section(section).Key(key).SetValue(val)
	err := cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Printf("Error saving config file: %s\n", err)
		return
	}
}
