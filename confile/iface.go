// INTERFACE APPDATA STRUCT-CLASS FROM SRC/APPDATA

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

	GetCellList() map[string][]src.CellData
	SetCellList(map[string][]src.CellData)
	SetCellListAppend(src.CellData, string)
	SetDeleteCell(int, string)

	GetFilepath() string
	SetFilepath(string)

	GetBar() *widget.ProgressBar
	SetBar(*widget.ProgressBar)

	GetTicker() *time.Ticker
	SetTicker(*time.Ticker)

	GetMainBar() *widget.ProgressBar
	SetMainBar(*widget.ProgressBar)

	GetInfoDialog() *src.InfoDialog
}
