package confile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func SaveFile(NewAppData *src.AppData) {
	code, err := json.Marshal(NewAppData.CellList)
	if err != nil {
		fmt.Println("Error", err)
	}
	if GetFromIni("file", "path") == "" {
		createNewFile(NewAppData, code)
	} else {
		saveInFile(NewAppData, code)
	}
}

func createNewFile(NewAppData *src.AppData, code []byte) {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				NewAppData.SetFilepath(uc.URI().Path())
				SaveToIni("file", "path", NewAppData.GetFilepath())
				io.WriteString(uc, string(code))
				NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), CreateList(NewAppData)))
			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)
}

func createNewRotorFile(NewAppData *src.AppData, code []byte) {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				NewAppData.SetFilepath(uc.URI().Path())
				/* SaveToIni(NewAppData.GetFilepath()) */
				io.WriteString(uc, string(code))
				NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), CreateList(NewAppData)))
			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)
}

func saveInFile(NewAppData *src.AppData, code []byte) {
	file, err := os.Open(GetFromIni("file", "path"))
	defer file.Close()
	if err != nil {
		fmt.Printf("1Error opening file: %s\n", err)
		NewAppData.SetFilepath("")
		SaveToIni("file", "path", NewAppData.GetFilepath())
		dialog.ShowCustom("Not File", "", widget.NewLabel("File not found"), NewAppData.GetWindow())
		return
	} else {
		ioutil.WriteFile(GetFromIni("file", "path"), code, 0644)
		NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), CreateList(NewAppData)))
	}
}
