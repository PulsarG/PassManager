// HANDLER CONFIG FILE
// READ, WHRITE, TAKE

package confile

import (
	// "PassManager/errs"
	// "fmt"

	"github.com/go-ini/ini"
)

func cfgHandler() *ini.File {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		ErrorLog(err)
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
		ErrorLog(err)
		return
	}
}
