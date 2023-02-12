package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

type AppData struct {
	app        fyne.App
	mainWindow fyne.Window
	canvas     fyne.Canvas

	CellList  []CellData
	entryCode widget.Entry

	filePath       string
	controlLenList int

	copySec float64

	timeBar  *widget.ProgressBar
	timeTick *time.Ticker
}

func NewAppData(a fyne.App, w fyne.Window, c fyne.Canvas) *AppData {
	return &AppData{
		app:        a,
		mainWindow: w,
		canvas:     c,
		copySec:    10.0,
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

func (a *AppData) GetControlLenList() int {
	return a.controlLenList
}

/* func (a *AppData) GetCellList() []CellData {
	return a.cellList
}

func (a *AppData) SetCellList(list []CellData) {
	a.cellList = list
} */

func (a *AppData) SetFilepath(s string) {
	a.filePath = s
}

func (a *AppData) SetControlLen(i int) {
	a.controlLenList = i
}

func (a *AppData) GetBar() *widget.ProgressBar {
	return a.timeBar
}

func (a *AppData) GetTicker() *time.Ticker {
	return a.timeTick
}

func (a *AppData) SetBar(b *widget.ProgressBar) {
	a.timeBar = b
}

func (a *AppData) SetTicker(t *time.Ticker) {
	a.timeTick = t
}

func (a *AppData) GetCopysec() float64 {
	return a.copySec
}

func (a *AppData) SetCopysec(f float64) {
	a.copySec = f
}
