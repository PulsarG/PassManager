package confile

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "io/ioutil"
	// "os"

	// "PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/dialog"
	// "fyne.io/fyne/v2/widget"
	// "github.com/PulsarG/Enigma"
)

func BuildList(iface InfaceApp) {
	a := CreateMangerBtns(iface)
	a.Resize(fyne.NewSize(150, 400))
	iface.GetCanvas().SetContent(container.NewHBox(a, CreateList(iface)))
}
