package confile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	/* "PassManager/src" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

/* type InfaceApp interface {
	GetWindow() fyne.Window
	GetCanvas() fyne.Canvas
	SetCopysec(int)
	GetCellList() []src.CellData
	GetFilepath() string
	SetFilepath(string)
} */

func SaveFile(iface InfaceApp) {
	code, err := json.Marshal(iface.GetCellList())
	if err != nil {
		fmt.Println("Error", err)
	}
	if GetFromIni("file", "path") == "" {
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
				SaveToIni("file", "path", iface.GetFilepath())
				io.WriteString(uc, string(code))
				iface.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(iface), CreateList(iface)))
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
				/* SaveToIni(NewAppData.GetFilepath()) */
				io.WriteString(uc, string(code))
				iface.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(iface), CreateList(iface)))
			} else {
				return
			}
		}, iface.GetWindow(),
	)
}

func saveInFile(iface InfaceApp, code []byte) {
	file, err := os.Open(GetFromIni("file", "path"))
	defer file.Close()
	if err != nil {
		fmt.Printf("1Error opening file: %s\n", err)
		iface.SetFilepath("")
		SaveToIni("file", "path", iface.GetFilepath())
		dialog.ShowCustom("Not File", "", widget.NewLabel("File not found"), iface.GetWindow())
		return
	} else {
		ioutil.WriteFile(GetFromIni("file", "path"), code, 0644)
		iface.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(iface), CreateList(iface)))
	}
}
