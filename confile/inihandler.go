package confile

import (
	/* "encoding/json" */
	"fmt"
	/* "image/color"
	"io"
	"io/ioutil"
	"os"
	*/
	/* "PassManager/cons"
	"PassManager/elem"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma" */
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
