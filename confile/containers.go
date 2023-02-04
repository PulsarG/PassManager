package confile

import (
	/* "encoding/json"
	"fmt" */
	"image/color"
	/* "io"
	"io/ioutil"
	"os" */

	"PassManager/cons"
	"PassManager/elem"
	"PassManager/src"

	"fyne.io/fyne/v2"
	/* "fyne.io/fyne/v2/app" */
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma"
	/* 	"github.com/go-ini/ini" */)

func CreateMangerBtns(NewAppData *src.AppData) *fyne.Container {
	NewAppData.GetEntryCode().PlaceHolder = cons.ENTER_KEY_PLACEHOLDER

	btnCreateCell := createColorBtn(cons.BTN_LABEL_CREATE_NEW_CELL, NewAppData, func() { createNewCellList(NewAppData) })

	containerAddandKey := container.NewGridWithColumns(2, btnCreateCell, NewAppData.GetEntryCode())

	btnOpen := createColorBtn(cons.BTN_LABEL_OPEN, NewAppData, func() { OpenFile(NewAppData) })
	/* btnSave := createColorBtn(cons.BTN_LABEL_SAVE, NewAppData, func() { saveFile(NewAppData) }) */
	containerOpenSaveBtn := container.NewGridWithColumns(1, btnOpen)

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
	if NewAppData.GetEntryCode().Text != "" {
		newCell := src.CreateNewCell()
		form := widget.NewForm(
			widget.NewFormItem(cons.FORM_LABEL_NAME, newCell.GetLabel()),
			widget.NewFormItem(cons.FORM_LABEL_LOGIN, newCell.GetLogin()),
			widget.NewFormItem(cons.FORM_LABEL_PASS, newCell.GetPass()),
		)
		comt := container.NewVBox(form, elem.NewButton("random pass", func() {}))
		dialog.ShowCustomConfirm(cons.DIALOG_CREATE_CELL_NAME,
			"Add",
			"Close",
			comt, func(b bool) {
				if b {
					setDataFromDialogCell(newCell, NewAppData)
				} else {
					return
				}
			},
			NewAppData.GetWindow())
	} else {
		dialog.ShowCustom("Oops", "Ok", widget.NewLabel("Please entry Key-Word"), NewAppData.GetWindow())
	}
}

func setDataFromDialogCell(newCell *src.Cell, NewAppData *src.AppData) {
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

	NewAppData.GetCanvas().SetContent(container.NewVSplit(CreateMangerBtns(NewAppData), elem.CreateList(NewAppData)))

	SaveFile(NewAppData)
}
