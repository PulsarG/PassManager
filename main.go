package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"io/ioutil"
	"os"

	"PassManager/cell"
	"PassManager/cons"
	"PassManager/elem"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma"
	"github.com/go-ini/ini"
)

var filePath string
var isSave bool

func main() {
	isSave = false
	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME)
	canvas := mainWindow.Canvas()

	NewAppData := src.NewAppData(App, mainWindow, canvas)

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	if isSave {
		openFile(NewAppData)
	} else {
		canvas.SetContent(container.NewCenter(createMangerBtns(NewAppData)))
	}

	mainWindow.Show()
	App.Run()
}

func createMangerBtns(NewAppData *src.AppData) *fyne.Container {
	NewAppData.GetEntryCode().PlaceHolder = cons.ENTER_KEY_PLACEHOLDER

	btnCreateCell := createColorBtn(cons.BTN_LABEL_CREATE_NEW_CELL, NewAppData, func() { createNewCellList(NewAppData) })

	containerAddandKey := container.NewGridWithColumns(2, btnCreateCell, NewAppData.GetEntryCode())

	btnOpen := createColorBtn(cons.BTN_LABEL_OPEN, NewAppData, func() { openFile(NewAppData) })
	btnSave := createColorBtn(cons.BTN_LABEL_SAVE, NewAppData, func() { saveFile(NewAppData) })
	containerOpenSaveBtn := container.NewGridWithColumns(2, btnOpen, btnSave)

	btnOpenCustomRotor := createColorBtn(cons.BTN_LABEL_OPEN_ROTOR, NewAppData, func() {})
	btnCreateCustomRotor := createColorBtn(cons.BTN_LABEL_CREATE_CUSTOM_ROTOR, NewAppData, func() {})
	containerCustomRotor := container.NewGridWithColumns(2, btnOpenCustomRotor, btnCreateCustomRotor)

	containerManager := container.NewGridWithRows(3, containerAddandKey, containerOpenSaveBtn, containerCustomRotor)
	return containerManager
}

func createColorBtn(label string, NewAppData *src.AppData, f func()) *fyne.Container {
	return container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color.NRGBA{R: 11, G: 78, B: 150, A: 1}),
		elem.NewButton(label, f),
	)
}

func createNewCellList(NewAppData *src.AppData) {
	newCell := cell.CreateNewCell()

	form := widget.NewForm(
		widget.NewFormItem(cons.FORM_LABEL_NAME, newCell.GetLabel()),
		widget.NewFormItem(cons.FORM_LABEL_LOGIN, newCell.GetLogin()),
		widget.NewFormItem(cons.FORM_LABEL_PASS, newCell.GetPass()),
	)

	form.OnSubmit = func() {
		setDataFromDialogCell(newCell, NewAppData)
	}
	form.SubmitText = "Save New Data"
	comt := container.NewVBox(form, elem.NewButton("random pass", func() {}))

	dialog.ShowCustom(cons.DIALOG_CREATE_CELL_NAME, "Close", comt, NewAppData.GetWindow())
}

func setDataFromDialogCell(newCell *cell.Cell, NewAppData *src.AppData) {
	newCellData := src.NewCellData()
	var err bool

	newCellData.Label = newCell.GetLabel().Text
	newCellData.Login, err = enigma.StartCrypt(newCell.GetLogin().Text, NewAppData.GetEntryCode().Text)
	if !err {
		dialog.ShowCustom("Error", "OK", widget.NewLabel(newCellData.Login), NewAppData.GetWindow())
		return
	}
	newCellData.Pass, err = enigma.StartCrypt(newCell.GetPass().Text, NewAppData.GetEntryCode().Text)
	if !err {
		dialog.ShowCustom("Error", "OK", widget.NewLabel(newCellData.Pass), NewAppData.GetWindow())
		return
	}

	NewAppData.CellList = append(NewAppData.CellList, *newCellData)

	NewAppData.GetCanvas().SetContent(container.NewVSplit(createMangerBtns(NewAppData), elem.CreateList(NewAppData)))

	saveFile(NewAppData)
}

func saveFile(NewAppData *src.AppData) {
	code, err := json.Marshal(NewAppData.CellList)
	if err != nil {
		fmt.Println("Error", err)
	}

	if getFilepathFromIni() == "" {
		createNewFile(NewAppData, code)
	} else {
		saveInFile(NewAppData, code)
	}

}

func saveInFile(NewAppData *src.AppData, code []byte) {
	file, err := os.Open(getFilepathFromIni())
	defer file.Close()
	if err != nil {
		fmt.Printf("1Error opening file: %s\n", err)
		filePath = ""
		firstSaveIni()
		dialog.ShowCustom("Not File", "", widget.NewLabel("123"), NewAppData.GetWindow())
		return
	} else {
		ioutil.WriteFile(getFilepathFromIni(), code, 0644)
	}
}

func createNewFile(NewAppData *src.AppData, code []byte) {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				filePath = uc.URI().Path()
				firstSaveIni()
				io.WriteString(uc, string(code))
				NewAppData.GetCanvas().SetContent(container.NewVBox(createMangerBtns(NewAppData), elem.CreateList(NewAppData)))
			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)

}

func getFilepathFromIni() string {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return ""
	}
	return cfg.Section("file").Key("path").String()
}

func findFile(NewAppData *src.AppData) {
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {

				filePath = uc.URI().Path()
				firstSaveIni()

				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewAppData.CellList)
				if err != nil {
					panic(err)
				}

				NewAppData.GetCanvas().SetContent(container.NewVBox(createMangerBtns(NewAppData), elem.CreateList(NewAppData)))

			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)
}

func getDatafromFile(NewAppData *src.AppData) {
	file, err := os.Open(getFilepathFromIni())
	if err != nil {
		fmt.Printf("2Error opening file: %s\n", err)
		filePath = ""
		firstSaveIni()
		dialog.ShowCustom("Not File", "", widget.NewLabel("123"), NewAppData.GetWindow())
		return
	} else {
		result, _ := ioutil.ReadAll(file)
		err := json.Unmarshal(result, &NewAppData.CellList)
		if err != nil {
			panic(err)
		}
		NewAppData.GetCanvas().SetContent(container.NewVBox(createMangerBtns(NewAppData), elem.CreateList(NewAppData)))

	}
	defer file.Close()
}

func openFile(NewAppData *src.AppData) {
	if getFilepathFromIni() == "" {
		findFile(NewAppData)
	} else {
		getDatafromFile(NewAppData)
	}

}

func firstSaveIni() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Error loading config file: %s\n", err)
		return
	}

	cfg.Section("file").Key("path").SetValue(filePath)

	err = cfg.SaveTo("config.ini")
	if err != nil {
		fmt.Printf("Error saving config file: %s\n", err)
		return
	}
}
