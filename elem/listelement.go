package elem

import (
	"PassManager/cons"
	"fmt"
	"image/color"

	"PassManager/src"

	"fyne.io/fyne/v2"
	/* 	"fyne.io/fyne/v2/app" */
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	/* "fyne.io/fyne/v2/layout" */)

func createListElement(id int, label, login, pass string, NewAppData *src.AppData) *fyne.Container {

	txtBoundPass := binding.NewString()
	txtBoundPass.Set(pass)
	copyBtnPass := widget.NewButtonWithIcon(cons.BTN_LABEL_COPY_PASS, theme.ContentCopyIcon(), func() {
		if content, err := txtBoundPass.Get(); err == nil { // content - строка, которую надо расшифровать перед копированием или показом
			NewAppData.GetWindow().Clipboard().SetContent(content)
		}
		fmt.Println(NewAppData.GetEntryCode().Text)
	})

	txtBoundLogin := binding.NewString()
	txtBoundLogin.Set(login)
	copyBtnLogin := widget.NewButtonWithIcon(cons.BTN_LABEL_COPY_LOGIN, theme.ContentCopyIcon(), func() {
		if content, err := txtBoundLogin.Get(); err == nil {
			NewAppData.GetWindow().Clipboard().SetContent(content)
		}
		fmt.Println(NewAppData.GetEntryCode().Text)
	})

	contLabelandChek := container.NewGridWithColumns(3, NewButton("Edit", nil), NewButton("Delete", func() { deleteCell(id, NewAppData) }), NewButton(cons.BTN_LABEL_SHOW_LOGPASS, func() {
		showPass(copyBtnLogin, copyBtnPass, login, pass)
	}))
	line := canvas.NewLine(color.Black)
	contLogPass := container.NewGridWithColumns(2, copyBtnLogin, copyBtnPass)
	listElementContainer := container.NewVBox(line, widget.NewLabel(label), contLabelandChek, contLogPass)

	return listElementContainer
}

func showPass(copyBtnLogin *widget.Button, copyBtnPass *widget.Button, login, pass string) {
	if copyBtnLogin.Text == cons.BTN_LABEL_COPY_LOGIN {
		copyBtnPass.SetText(pass)
		copyBtnPass.Refresh()

		copyBtnLogin.SetText(login)
		copyBtnLogin.Refresh()
	} else {
		copyBtnPass.SetText(cons.BTN_LABEL_COPY_PASS)
		copyBtnPass.Refresh()

		copyBtnLogin.SetText(cons.BTN_LABEL_COPY_LOGIN)
		copyBtnLogin.Refresh()
	}
}

func CreateList(NewAppData *src.AppData) *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(NewAppData.CellList); i++ {
		containerListElement := createListElement(i, NewAppData.CellList[i].Label, NewAppData.CellList[i].Login, NewAppData.CellList[i].Pass, NewAppData)
		listContainer.Add(containerListElement)
	}
	return listContainer
}

func deleteCell(id int, NewAppData *src.AppData) {

	NewAppData.CellList = append(NewAppData.CellList[:id], NewAppData.CellList[id+1:]...)

}
