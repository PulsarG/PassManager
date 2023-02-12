package confile

import (
	"fmt"
	
	"github.com/go-ini/ini"
)

func GetFilepathFromIni() string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return ""
	}
	return cfg.Section("file").Key("path").String()
}

func firstSaveIni(path string) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return
	}

	cfg.Section("file").Key("path").SetValue(path)

	err = cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Printf("Error saving config file: %s\n", err)
		return
	}
}
