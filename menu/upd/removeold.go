// REMOVE OLD FILES AFTER UPDATE APP

package upd

import (
	"PassManager/confile"
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
)

type InfaceApp interface {
	GetWindow() fyne.Window
}

func CheckOld() {
	ticker := time.NewTicker(time.Second)
CHECK:
	for range ticker.C {
		if confile.GetFromIni("data", "old") == "" {
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
	err := os.Remove(confile.GetFromIni("data", "old"))
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		confile.SaveToIni("data", "old", "")
		return true
	}
}
