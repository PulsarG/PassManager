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

func findFile(NewAppData *src.AppData) bool {
	isFind := false
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				NewAppData.SetFilepath(uc.URI().Path())
				SaveToIni("file", "path", NewAppData.GetFilepath())
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewAppData.CellList)
				if err != nil {
					panic(err)
				}
				NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), CreateList(NewAppData)))
				isFind = true
			} else {
				isFind = false
			}
		}, NewAppData.GetWindow(),
	)
	return isFind
}

func GetDatafromFile(NewAppData *src.AppData) {
	file, err := os.Open(GetFromIni("file", "path"))
	if err != nil {
		fmt.Printf("2Error opening file: %s\n", err)
		NewAppData.SetFilepath("")
		SaveToIni("file", "path", NewAppData.GetFilepath())
		dialog.ShowCustom("Not File", "Ok", widget.NewLabel("File not found. Please create new file"), NewAppData.GetWindow())
		NewAppData.GetCanvas().SetContent(container.NewCenter(CreateMangerBtns(NewAppData)))
	} else {
		result, _ := ioutil.ReadAll(file)
		err := json.Unmarshal(result, &NewAppData.CellList)
		if err != nil {
			panic(err)
		}
		NewAppData.GetCanvas().SetContent(container.NewVBox(CreateMangerBtns(NewAppData), CreateList(NewAppData)))

	}
	defer file.Close()
}

func OpenFile(NewAppData *src.AppData) bool {
	/* if GetFilepathFromIni() == "" { */
	return findFile(NewAppData)
	/* } else {
		GetDatafromFile(NewAppData)
	} */
}
