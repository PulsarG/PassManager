package main

import (
	"PassManager/confile"
	"PassManager/cons"
	"PassManager/menu"
	"PassManager/menu/upd"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	/* 	"fyne.io/fyne/v2/theme" */)

func main() {
	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME + upd.GetVersion())
	canvas := mainWindow.Canvas()

	NewAppData := src.NewAppData(App, mainWindow, canvas, confile.GetCopysecIni())

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	if confile.GetFilepathFromIni() != "" {
		confile.GetDatafromFile(NewAppData)
	} else {
		canvas.SetContent(container.NewCenter(confile.CreateMangerBtns(NewAppData)))
	}

	// App.Settings().SetTheme(theme.DarkTheme())
	mainWindow.SetMainMenu(menu.GetMenu(NewAppData))
	mainWindow.Show()
	App.Run()
}
