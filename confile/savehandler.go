package confile

import (
	"encoding/json"
	"fmt"

	/* "image/color" */
	"io"
	"io/ioutil"
	"os"

	/* "PassManager/cons" */
	"PassManager/elem"
	"PassManager/src"

	"fyne.io/fyne/v2"
	/* enigma "github.com/PulsarG/Enigma" */
	/* "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas" */
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	/* "fyne.io/fyne/v2/layout" */
	"fyne.io/fyne/v2/widget"
	/* "github.com/PulsarG/Enigma" */ /* "github.com/go-ini/ini" */)

func SaveFile(NewAppData *src.AppData) {
	code, err := json.Marshal(NewAppData.CellList)
	if err != nil {
		fmt.Println("Error", err)
	}
	if GetFilepathFromIni() == "" {
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
				firstSaveIni(NewAppData.GetFilepath())
				io.WriteString(uc, string(code))
				NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), elem.CreateList(NewAppData)))
			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)

}

func saveInFile(NewAppData *src.AppData, code []byte) {
	file, err := os.Open(GetFilepathFromIni())
	defer file.Close()
	if err != nil {
		fmt.Printf("1Error opening file: %s\n", err)
		NewAppData.SetFilepath("")
		firstSaveIni(NewAppData.GetFilepath())
		dialog.ShowCustom("Not File", "", widget.NewLabel("123"), NewAppData.GetWindow())
		return
	} else {
		ioutil.WriteFile(GetFilepathFromIni(), code, 0644)
		NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), elem.CreateList(NewAppData)))
	}
}
