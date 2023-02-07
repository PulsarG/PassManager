package main

import (
/* 	"time" */

	"PassManager/confile"
	"PassManager/cons"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

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

	/* go cycleUpdateList(NewAppData) */
	mainWindow.Show()
	App.Run()
}

/* func cycleUpdateList(NewAppData *src.AppData) {
	controlLen := NewAppData.GetControlLenList()

	ticker := time.Tick(time.Second)
	for range ticker {
		if len(NewAppData.CellList) != controlLen {
			confile.SaveFile(NewAppData)
			controlLen = NewAppData.GetControlLenList()
		}
	}
} */
