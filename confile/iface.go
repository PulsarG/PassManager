package confile

import (
	"time"

	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type InfaceApp interface {
	GetApp() fyne.App
	GetWindow() fyne.Window
	GetCanvas() fyne.Canvas

	GetEntryCode() *widget.Entry

	GetCopysec() int
	SetCopysec(int)

	GetCellList() []src.CellData
	SetCellList([]src.CellData)
	SetCellListAppend(src.CellData)
	SetDeleteCell(int)

	GetFilepath() string
	SetFilepath(string)

	GetBar() *widget.ProgressBar
	SetBar(*widget.ProgressBar)

	GetTicker() *time.Ticker
	SetTicker(*time.Ticker)
}
