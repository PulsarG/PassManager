package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"

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
)

func main() {

	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME)
	canvas := mainWindow.Canvas()

	NewAppData := src.NewAppData(App, mainWindow, canvas)

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	canvas.SetContent(container.NewCenter(createMangerBtns(NewAppData)))
	mainWindow.Show()
	App.Run()
}

func createMangerBtns(NewAppData *src.AppData) *fyne.Container {
	NewAppData.GetEntryCode().PlaceHolder = cons.ENTER_KEY_PLACEHOLDER

	btnCreateCell := createColorBtn(cons.BTN_LABEL_CREATE_NEW_CELL, NewAppData, func() { createNewCellList(NewAppData) })
	btnOpen := createColorBtn(cons.BTN_LABEL_OPEN, NewAppData, func() { openFile(NewAppData) })
	btnSave := createColorBtn(cons.BTN_LABEL_SAVE, NewAppData, func() { saveFile(NewAppData) })

	containerAddandKey := container.NewGridWithColumns(2, btnCreateCell, NewAppData.GetEntryCode())

	containerOpenSaveBtn := container.NewGridWithColumns(2, btnOpen, btnSave)
	containerManager := container.NewGridWithRows(2, containerAddandKey, containerOpenSaveBtn)
	return containerManager
}

func createColorBtn(label string, NewAppData *src.AppData, f func()) *fyne.Container {
	btnCreate := elem.NewButton(label, f)
	color := color.RGBA{11, 78, 150, 1}
	btn := container.New(
		layout.NewMaxLayout(),
		btnCreate,
		canvas.NewRectangle(color),
	)
	return container.NewWithoutLayout(btn)
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

	dialog.ShowCustom(cons.DIALOG_CREATE_CELL_NAME, "Close", form, NewAppData.GetWindow())
}

func setDataFromDialogCell(newCell *cell.Cell, NewAppData *src.AppData) {
	newCellData := src.NewCellData()

	newCellData.Label = newCell.GetLabel().Text
	newCellData.Login = newCell.GetLogin().Text
	newCellData.Pass = newCell.GetPass().Text

	NewAppData.CellList = append(NewAppData.CellList, *newCellData)

	NewAppData.GetCanvas().SetContent(container.NewVSplit(createMangerBtns(NewAppData), elem.CreateList(NewAppData)))

	fmt.Println(NewAppData.CellList)
}

func saveFile(NewAppData *src.AppData) {
	code, err := json.Marshal(NewAppData.CellList)
	if err != nil {
		fmt.Println("Error", err)
	}

	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				io.WriteString(uc, string(code))
				NewAppData.GetCanvas().SetContent(container.NewVBox(createMangerBtns(NewAppData), elem.CreateList(NewAppData)))
			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)
}

func openFile(NewAppData *src.AppData) {
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
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
