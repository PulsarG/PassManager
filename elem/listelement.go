package elem

import (
	"fyne.io/fyne/v2"
	/* 	"fyne.io/fyne/v2/app" */
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateListElement(label, login, pass string) *fyne.Container {
	contLabelandChek := container.NewGridWithColumns(2, widget.NewLabel(label), widget.NewCheck("Show Log-Pass", nil))
	contLogPass := container.NewGridWithColumns(2, NewButton(login, nil), NewButton(pass, nil))
	listElementContainer := container.NewVBox(contLabelandChek, contLogPass)
	return listElementContainer
}
