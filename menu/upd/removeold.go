// REMOVE OLD FILES AFTER UPDATE APP

package upd

import (
	"PassManager/errloger"
	// "PassManager/confile"
	// "fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"

	"github.com/PulsarG/ConfigManager"
)

type InfaceApp interface {
	GetWindow() fyne.Window
}

func CheckOld() {
	ticker := time.NewTicker(time.Second)
CHECK:
	for range ticker.C {
		if inihandler.GetFromIni("data", "old") == "" {
			ticker.Stop()
			break CHECK
		} else {
			if removeOld() {
				ticker.Stop()
				break CHECK
			}
		}
	}
}

func removeOld() bool {
	err := os.Remove(inihandler.GetFromIni("data", "old"))
	if err != nil {
		errloger.ErrorLog(err)
		return false
	} else {
		inihandler.SaveToIni("data", "old", "")
		return true
	}
}
