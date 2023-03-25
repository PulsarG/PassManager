// READ PASS-BASE AND ROTOR

package confile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"PassManager/src"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/PulsarG/Enigma"
)

func findFile(iface InfaceApp) bool {
	var cellData map[string][]src.CellData
	isFind := false
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				iface.SetFilepath(uc.URI().Path())
				SaveToIni("file", "path", iface.GetFilepath())
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &cellData)
				if err != nil {
					panic(err)
				}
				iface.SetCellList(cellData)
				iface.GetCanvas().SetContent(container.NewHSplit(CreateMangerBtns(iface), CreateList(iface)))
				isFind = true
			} else {
				isFind = false
			} // end if
		}, iface.GetWindow(),
	)
	return isFind
}

func GetRotorFromFile(iface InfaceApp) {
	var NewRotor [162]int
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewRotor)
				if err != nil {
					panic(err)
				}
				enigma.SetCustomRotor(NewRotor)

			} // end if
		}, iface.GetWindow(),
	)
}

func GetDatafromFile(iface InfaceApp) {
	cellData := iface.GetCellList()
	file, err := os.Open(GetFromIni("file", "path"))
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		iface.SetFilepath("")
		SaveToIni("file", "path", iface.GetFilepath())
		dialog.ShowCustom("Not File", "Ok", widget.NewLabel("File not found. Please create new file"), iface.GetWindow())
		iface.GetCanvas().SetContent(container.NewCenter(CreateMangerBtns(iface)))
	} else {
		result, _ := ioutil.ReadAll(file)
		err := json.Unmarshal(result, &cellData)
		if err != nil { // if inner
			panic(err)
		} // end if inner
		iface.SetCellList(cellData)

		iface.GetCanvas().SetContent(container.NewHSplit(CreateMangerBtns(iface), CreateList(iface)))
	} // end if
	defer file.Close()
}

func OpenFile(iface InfaceApp) bool {
	return findFile(iface)
}
