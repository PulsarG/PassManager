package confile

import (
	"fmt"
	"strconv"

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

func SaveCopysecIni(i int) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return
	}
	dur := strconv.Itoa(i)
	cfg.Section("data").Key("duration").SetValue(dur)

	err = cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Printf("Error saving config file: %s\n", err)
		return
	}
}
func GetCopysecIni() int {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
	}
	dur, err := strconv.Atoi(cfg.Section("data").Key("duration").String())
	if err != nil {
		fmt.Printf("Error read duration from config file: %s\n", err)
	}
	return dur
}
