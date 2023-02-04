package src

import (
	/* "encoding/json"
	"fmt"
	"io"

	"PassManager/cell"
	"PassManager/cons"
	"PassManager/elem" */

	"fyne.io/fyne/v2"
	/* "fyne.io/fyne/v2/app" */
	/* "fyne.io/fyne/v2/canvas" */
	/* "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog" */
	"fyne.io/fyne/v2/widget"
)

/* type CellData struct {
	Label string
	Login string
	Pass  string
} */

type AppData struct {
	app        fyne.App
	mainWindow fyne.Window
	canvas     fyne.Canvas

	CellList  []CellData
	entryCode widget.Entry

	filePath string
}

func NewAppData(a fyne.App, w fyne.Window, c fyne.Canvas) *AppData {
	return &AppData{
		app:        a,
		mainWindow: w,
		canvas:     c,
	}
}

func (a *AppData) GetApp() fyne.App {
	return a.app
}

func (a *AppData) GetWindow() fyne.Window {
	return a.mainWindow
}

func (a *AppData) GetCanvas() fyne.Canvas {
	return a.canvas
}

func (a *AppData) GetEntryCode() *widget.Entry {
	return &a.entryCode
}

func (a *AppData) GetFilepath() string {
	return a.filePath
}

func (a *AppData) SetFilepath(s string) {
	a.filePath = s
}

/* func NewCellData() *CellData {
	return &CellData{}
} */

