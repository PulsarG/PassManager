package elem

import (
	"PassManager/cons"
	"fmt"

	"fyne.io/fyne/v2"
	/* 	"fyne.io/fyne/v2/app" */
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	/* "fyne.io/fyne/v2/layout" */)

func CreateListElement(label, login, pass string, w fyne.Window, code string) *fyne.Container {

	txtBoundPass := binding.NewString()
	txtBoundPass.Set(pass)
	copyBtnPass := widget.NewButtonWithIcon(cons.BTN_LABEL_COPY_PASS, theme.ContentCopyIcon(), func() {
		if content, err := txtBoundPass.Get(); err == nil { // content - строка, которую надо расшифровать перед копированием или показом
			w.Clipboard().SetContent(content)
		}
		fmt.Println(code)
	})

	txtBoundLogin := binding.NewString()
	txtBoundLogin.Set(login)
	copyBtnLogin := widget.NewButtonWithIcon(cons.BTN_LABEL_COPY_LOGIN, theme.ContentCopyIcon(), func() {
		if content, err := txtBoundLogin.Get(); err == nil {
			w.Clipboard().SetContent(content)
		}
		fmt.Println(code)
	})

	contLabelandChek := container.NewGridWithColumns(2, widget.NewLabel(label), NewButton(cons.BTN_LABEL_SHOW_LOGPASS, func() {
		showPass(copyBtnLogin, copyBtnPass, login, pass)
	}))

	contLogPass := container.NewGridWithColumns(2, copyBtnLogin, copyBtnPass)
	listElementContainer := container.NewVBox(contLabelandChek, contLogPass)
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
