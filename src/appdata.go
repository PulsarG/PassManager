package src

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type AppData struct {
	app        fyne.App
	mainWindow fyne.Window
	canvas     fyne.Canvas

	cellList  map[string][]CellData
	entryCode widget.Entry

	filePath       string
	controlLenList int

	copySec int

	timeBar  *widget.ProgressBar
	timeTick *time.Ticker

	mainBar *widget.ProgressBar
}

func NewAppData(a fyne.App, w fyne.Window, c fyne.Canvas, i int) *AppData {
	return &AppData{
		app:        a,
		mainWindow: w,
		canvas:     c,
		copySec:    i,

		cellList: make(map[string][]CellData),

		mainBar: createMainBar(),
	}
}

func createMainBar() *widget.ProgressBar {
	b := widget.NewProgressBar()
	b.Hide()
	return b
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

// File path
func (a *AppData) GetFilepath() string {
	return a.filePath
}

func (a *AppData) SetFilepath(s string) {
	a.filePath = s
}

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

// CopySec
func (a *AppData) GetCopysec() int {
	return a.copySec
}

func (a *AppData) SetCopysec(i int) {
	a.copySec = i
}

// Cell List
func (a *AppData) GetCellList() map[string][]CellData {
	return a.cellList
}

func (a *AppData) SetCellListAppend(newCellData CellData, s string) {
	if CL, ok := a.cellList[s]; ok {
		CL = append(CL, newCellData)
		a.cellList[s] = CL
	} else {
		var newCL []CellData
		newCL = append(newCL, newCellData)
		a.cellList[s] = newCL
	}
}

func (a *AppData) SetDeleteCell(id int, s string) {
	a.cellList[s] = append(a.cellList[s][:id], a.cellList[s][id+1:]...)
}

func (a *AppData) SetCellList(list map[string][]CellData) {
	a.cellList = list
}

func (a *AppData) GetMainBar() *widget.ProgressBar {
	return a.mainBar
}
func (a *AppData) SetMainBar(b *widget.ProgressBar) {
	a.mainBar = b
}
