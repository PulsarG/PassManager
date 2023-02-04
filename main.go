package main

import (
	/* 	"encoding/json"
	   	"fmt"
	   	"image/color"
	   	"io"
	   	"io/ioutil"
	   	"os" */

	/* 	"PassManager/cell" */
	"PassManager/confile"
	"PassManager/cons"
	/* 	"PassManager/elem" */
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	/* "fyne.io/fyne/v2/canvas" */
	"fyne.io/fyne/v2/container"
	/* "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma" *//* 	"github.com/go-ini/ini" */)

func main() {
	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME)
	canvas := mainWindow.Canvas()

	NewAppData := src.NewAppData(App, mainWindow, canvas)

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	if confile.GetFilepathFromIni() != "" {
		confile.GetDatafromFile(NewAppData)
	} else {
		canvas.SetContent(container.NewCenter(confile.CreateMangerBtns(NewAppData)))
	}

	mainWindow.Show()
	App.Run()
}

/* func setDataFromDialogCell(newCell *src.Cell, NewAppData *src.AppData) {
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

	NewAppData.GetCanvas().SetContent(container.NewVSplit(elem.CreateMangerBtns(NewAppData), elem.CreateList(NewAppData)))

	conf.SaveFile(NewAppData)
} */
