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
		if _, err := os.Stat(confile.GetFromIni("data", "old")); os.IsNotExist(err) {
			ticker.Stop()
			break CHECK
		} else {
			removeOld()
			confile.SaveToIni("data", "old", "")
			ticker.Stop()
			break CHECK
		}
	}
}

func removeOld() {
	err := os.Remove(confile.GetFromIni("data", "old"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
