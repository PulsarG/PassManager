package src

import (
	/* "PassManager/confile" */
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AppData struct {
	app        fyne.App
	mainWindow fyne.Window
	canvas     fyne.Canvas

	cellList  []CellData
	entryCode widget.Entry

	filePath       string
	controlLenList int

	copySec int

	timeBar  *widget.ProgressBar
	timeTick *time.Ticker
}

func NewAppData(a fyne.App, w fyne.Window, c fyne.Canvas, i int) *AppData {
	return &AppData{
		app:        a,
		mainWindow: w,
		canvas:     c,
		copySec:    i,
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

func (a *AppData) GetControlLenList() int {
	return a.controlLenList
}

//
// File path
func (a *AppData) GetFilepath() string {
	return a.filePath
}

func (a *AppData) SetFilepath(s string) {
	a.filePath = s
}

//
// Ticker and Bar
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

//
// CopySec 
func (a *AppData) GetCopysec() int {
	return a.copySec
}

func (a *AppData) SetCopysec(i int) {
	a.copySec = i
}

//
// Cell List
func (a *AppData) GetCellList() []CellData {
	return a.cellList
}

func (a *AppData) SetCellListAppend(newCellData CellData) {
	a.cellList = append(a.cellList, newCellData)
}

func (a *AppData) SetDeleteCell(id int) {
	a.cellList = append(a.cellList[:id], a.cellList[id+1:]...)
}

func (a *AppData) SetCellList(list []CellData) {
	a.cellList = list
}
