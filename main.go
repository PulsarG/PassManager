package main

import (
	"PassManager/confile"
	"PassManager/cons"
	"PassManager/menu"
	"PassManager/menu/upd"
	"PassManager/src"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	go upd.CheckOld()

	mainApp := app.New()
	mainWindow := mainApp.NewWindow(cons.WINDOW_NAME + confile.GetFromIni("data", "version"))
	duration, _ := strconv.Atoi(confile.GetFromIni("data", "duration"))
	NewAppData := src.NewAppData(mainApp, mainWindow, mainWindow.Canvas(), duration)

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	selectWindowContent(NewAppData)

	mainWindow.SetMainMenu(menu.GetMenu(NewAppData))

	mainWindow.CenterOnScreen()
	mainWindow.Show()
	mainApp.Run()
}

func selectWindowContent(NewAppData *src.AppData) {
	if confile.GetFromIni("file", "path") != "" {
		confile.GetDatafromFile(NewAppData)
	} else {
		NewAppData.GetCanvas().SetContent(container.NewCenter(confile.CreateMangerBtns(NewAppData)))
	} // end if
}
