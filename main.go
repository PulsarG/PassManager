package main

import (
	"PassManager/confile"
	"PassManager/cons"
	"PassManager/menu"
	"PassManager/src"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	/* 	"fyne.io/fyne/v2/theme" */)

func main() {
	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME + confile.GetFromIni("data", "version"))
	duration, _ := strconv.Atoi(confile.GetFromIni("data", "duration"))
	NewAppData := src.NewAppData(App, mainWindow, mainWindow.Canvas(), duration)
	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	selectWindowContent(NewAppData)

	// App.Settings().SetTheme(theme.DarkTheme())
	mainWindow.SetMainMenu(menu.GetMenu(NewAppData))
	mainWindow.Show()
	App.Run()
}

func selectWindowContent(NewAppData *src.AppData) {
	if confile.GetFromIni("file", "path") != "" {
		confile.GetDatafromFile(NewAppData)
	} else {
		NewAppData.GetCanvas().SetContent(container.NewCenter(confile.CreateMangerBtns(NewAppData)))
	}
}
