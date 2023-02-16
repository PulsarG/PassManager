package confile

import (
	"PassManager/cons"
	"fmt"
	"image/color"
	"time"

	/* "PassManager/confile" */
	"PassManager/elem"
	/* "PassManager/src" */

	"fyne.io/fyne/v2"
	/* 	"fyne.io/fyne/v2/app" */
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma"
)

func createListElement(id int, label, login, pass string, iface InfaceApp) *fyne.Container {
	barCopy := widget.NewProgressBar()
	barCopy.Hide()
	copyBtnPass := createBtnWithIcon(iface, pass, cons.BTN_LABEL_COPY_PASS, barCopy)
	copyBtnLogin := createBtnWithIcon(iface, login, cons.BTN_LABEL_COPY_LOGIN, barCopy)

	contManageCell := container.NewGridWithColumns(3,
		elem.NewButton("Edit", func() { editCellDialog(iface, id) }),
		elem.NewButton("Delete",
			func() { deleteCell(id, iface) }),
		elem.NewButton(cons.BTN_LABEL_SHOW_LOGPASS,
			func() {
				showPass(iface, copyBtnLogin, copyBtnPass, login, pass)
			}))

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 5

	nameLabel := widget.NewLabel(label)
	nameLabel.TextStyle.Bold = true
	nameLabel.TextStyle.Italic = true
	contNameLogPass := container.NewGridWithColumns(3,
		/* container.NewCenter(container.New(
		layout.NewMaxLayout(), */
		nameLabel,
		/* canvas.NewRectangle(color.RGBA{17, 0, 123, 1}),
		)), */
		copyBtnLogin,
		copyBtnPass,
	)

	listElementContainer := container.NewVBox(line, contManageCell, contNameLogPass, barCopy)

	if id%2 != 0 {
		return listElementContainer
	} else {
		color := color.RGBA{0, 0, 180, 1}
		listElementContainerColor := container.New(
			layout.NewMaxLayout(),
			listElementContainer,
			canvas.NewRectangle(color),
		)
		return listElementContainerColor
	}
}

func createBtnWithIcon(iface InfaceApp, data, name string, barCopy *widget.ProgressBar) *widget.Button {
	txtBoundPass := binding.NewString()
	txtBoundPass.Set(data)
	copyBtn := widget.NewButtonWithIcon(name, theme.ContentCopyIcon(), func() {
		if iface.GetTicker() != nil && iface.GetBar() != nil {
			iface.GetBar().Hide()
			iface.GetTicker().Stop()
		}
		iface.SetBar(barCopy)
		go copyAndBarr(txtBoundPass, iface, barCopy)
	})
	return copyBtn
}

func copyAndBarr(txtBoundPass binding.String, iface InfaceApp, barCopy *widget.ProgressBar) {
	content, err := txtBoundPass.Get()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	toCopy, errEnc := enigma.StartCrypt(content, iface.GetEntryCode().Text)
	if !errEnc {
		dialog.ShowCustom("Error", "OK", widget.NewLabel(toCopy), iface.GetWindow())
		return
	}

	iface.GetWindow().Clipboard().SetContent(toCopy)
	progressBarLine(iface, barCopy)
	<-time.After(time.Duration(iface.GetCopysec()) * time.Second)
	iface.GetWindow().Clipboard().SetContent("")
}

func progressBarLine(iface InfaceApp, barCopy *widget.ProgressBar) {
	timeSecond := float64(iface.GetCopysec())
	barCopy.Value = timeSecond
	barCopy.Min = 0.0
	barCopy.Max = timeSecond
	barCopy.Show()

	/* ticker := time.NewTicker(time.Second) */
	iface.SetTicker(time.NewTicker(time.Second))
	for range iface.GetTicker().C {
		timeSecond--
		fmt.Println("Left ", timeSecond)
		barCopy.SetValue(timeSecond)
		if timeSecond == 0.0 {
			barCopy.Hide()
			iface.GetTicker().Stop()
		}
	}
}

func showPass(iface InfaceApp, copyBtnLogin *widget.Button, copyBtnPass *widget.Button, login, pass string) {
	if copyBtnLogin.Text == cons.BTN_LABEL_COPY_LOGIN {
		openPass, _ := enigma.StartCrypt(pass, iface.GetEntryCode().Text)

		copyBtnPass.SetText(openPass)
		copyBtnPass.Refresh()
		openLogin, _ := enigma.StartCrypt(login, iface.GetEntryCode().Text)
		copyBtnLogin.SetText(openLogin)
		copyBtnLogin.Refresh()
	} else {
		copyBtnPass.SetText(cons.BTN_LABEL_COPY_PASS)
		copyBtnPass.Refresh()

		copyBtnLogin.SetText(cons.BTN_LABEL_COPY_LOGIN)
		copyBtnLogin.Refresh()
	}
}

func CreateList(iface InfaceApp) *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(iface.GetCellList()); i++ {
		containerListElement := createListElement(i, iface.GetCellList()[i].Label, iface.GetCellList()[i].Login, iface.GetCellList()[i].Pass, iface)
		listContainer.Add(containerListElement)
	}
	return listContainer
}

func deleteCell(id int, iface InfaceApp) {
	dialog.ShowConfirm("DELETE?", "REALY?", func(b bool) {
		if b {
			iface.SetDeleteCell(id)
			/* iface.SetControlLen(len(iface.CellList)) */
			SaveFile(iface)
		}
	}, iface.GetWindow())
}

func editCellDialog(iface InfaceApp, id int) {
	if iface.GetEntryCode().Text == "" {
		dialog.ShowInformation("Opps", "Please enter key-word", iface.GetWindow())
		return
	} else {
		var newData [3]widget.Entry
		newData[0].PlaceHolder = "New Label"
		newData[1].PlaceHolder = "New Login"
		newData[2].PlaceHolder = "New Password"
		forms := container.NewVBox(&newData[0], &newData[1], &newData[2])
		dialog.ShowCustomConfirm("Edit", "Accept", "Exit", forms, func(b bool) {
			if b {
				editCell(id, newData, iface)
			}
		}, iface.GetWindow())
	}
}

func editCell(id int, newData [3]widget.Entry, iface InfaceApp) {
	if newData[0].Text != "" {
		iface.GetCellList()[id].Label = newData[0].Text
	}
	if newData[1].Text != "" {
		s, b := enigma.StartCrypt(newData[1].Text, iface.GetEntryCode().Text)
		if !b {
			return
		}
		iface.GetCellList()[id].Login = s
	}
	if newData[2].Text != "" {
		s, b := enigma.StartCrypt(newData[2].Text, iface.GetEntryCode().Text)
		if !b {
			return
		}
		iface.GetCellList()[id].Pass = s
	}
	SaveFile(iface)
}

func popUpMenu(iface InfaceApp) *widget.PopUpMenu {
	popMenu := fyne.NewMenu("123", fyne.NewMenuItem("321", func() {}))
	pop := widget.NewPopUpMenu(popMenu, iface.GetCanvas())
	return pop
}
