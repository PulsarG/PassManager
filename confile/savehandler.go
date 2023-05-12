// WRITE PASS-BASE AND ROTOR IN FILE

package confile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"PassManager/errloger"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/ConfigManager"
)

func SaveFile(iface InfaceApp) {
	code, err := json.Marshal(iface.GetCellList())
	if err != nil {
		errloger.ErrorLog(err)
	}
	if inihandler.GetFromIni("file", "path") == "" {
		createNewFile(iface, code)
	} else {
		saveInFile(iface, code)
	}
}

func createNewFile(iface InfaceApp, code []byte) {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				iface.SetFilepath(uc.URI().Path())
				inihandler.SaveToIni("file", "path", iface.GetFilepath())
				io.WriteString(uc, string(code))
				// ***
				BuildList(iface)
				// a := CreateMangerBtns(iface)
				// a.Resize(fyne.NewSize(150, 400))
				// iface.GetCanvas().SetContent(container.NewHBox(a, CreateList(iface)))
				// iface.GetCanvas().SetContent(container.NewHSplit(CreateMangerBtns(iface), CreateList(iface)))
			} else {
				return
			}
		}, iface.GetWindow(),
	)
}

func createNewRotorFile(iface InfaceApp, code []byte) {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				iface.SetFilepath(uc.URI().Path())
				io.WriteString(uc, string(code))
				// ***
				BuildList(iface)
				// a := CreateMangerBtns(iface)
				// a.Resize(fyne.NewSize(150, 400))
				// iface.GetCanvas().SetContent(container.NewHBox(a, CreateList(iface)))
				// iface.GetCanvas().SetContent(container.NewHSplit(CreateMangerBtns(iface), CreateList(iface)))
			} else {
				return
			}
		}, iface.GetWindow(),
	)
}

func saveInFile(iface InfaceApp, code []byte) {
	file, err := os.Open(inihandler.GetFromIni("file", "path"))
	defer file.Close()
	if err != nil {
		fmt.Printf("1Error opening file: %s\n", err)
		iface.SetFilepath("")
		inihandler.SaveToIni("file", "path", iface.GetFilepath())
		dialog.ShowCustom("Not File", "", widget.NewLabel("File not found"), iface.GetWindow())
		return
	} else {
		ioutil.WriteFile(inihandler.GetFromIni("file", "path"), code, 0644)
		// ***
		BuildList(iface)
		// a := CreateMangerBtns(iface)
		// a.Resize(fyne.NewSize(150, 400))
		// iface.GetCanvas().SetContent(container.NewHBox(a, CreateList(iface)))
		// iface.GetCanvas().SetContent(container.NewHSplit(CreateMangerBtns(iface), CreateList(iface)))
	}
}
