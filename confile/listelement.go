package confile

import (
	"PassManager/cons"
	"fmt"
	"image/color"
	"time"

	/* "PassManager/confile" */
	"PassManager/elem"
	"PassManager/src"

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

func createListElement(id int, label, login, pass string, NewAppData *src.AppData) *fyne.Container {
	barCopy := widget.NewProgressBar()
	barCopy.Hide()
	copyBtnPass := createBtnWithIcon(NewAppData, pass, cons.BTN_LABEL_COPY_PASS, barCopy)
	copyBtnLogin := createBtnWithIcon(NewAppData, login, cons.BTN_LABEL_COPY_LOGIN, barCopy)

	contManageCell := container.NewGridWithColumns(3,
		elem.NewButton("Edit", func() { editCellDialog(NewAppData, id) }),
		elem.NewButton("Delete",
			func() { deleteCell(id, NewAppData) }),
		elem.NewButton(cons.BTN_LABEL_SHOW_LOGPASS,
			func() {
				showPass(NewAppData, copyBtnLogin, copyBtnPass, login, pass)
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

func createBtnWithIcon(NewAppData *src.AppData, data, name string, barCopy *widget.ProgressBar) *widget.Button {
	txtBoundPass := binding.NewString()
	txtBoundPass.Set(data)
	copyBtn := widget.NewButtonWithIcon(name, theme.ContentCopyIcon(), func() {
		if NewAppData.GetTicker() != nil && NewAppData.GetBar() != nil {
			NewAppData.GetBar().Hide()
			NewAppData.GetTicker().Stop()
		}
		NewAppData.SetBar(barCopy)
		go copyAndBarr(txtBoundPass, NewAppData, barCopy)
	})
	return copyBtn
}

func copyAndBarr(txtBoundPass binding.String, NewAppData *src.AppData, barCopy *widget.ProgressBar) {
	content, err := txtBoundPass.Get()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	toCopy, errEnc := enigma.StartCrypt(content, NewAppData.GetEntryCode().Text)
	if !errEnc {
		dialog.ShowCustom("Error", "OK", widget.NewLabel(toCopy), NewAppData.GetWindow())
		return
	}

	NewAppData.GetWindow().Clipboard().SetContent(toCopy)
	progressBarLine(NewAppData, barCopy)
	<-time.After(time.Duration(NewAppData.GetCopysec()) * time.Second)
	NewAppData.GetWindow().Clipboard().SetContent("")
}

func progressBarLine(NewAppData *src.AppData, barCopy *widget.ProgressBar) {
	timeSecond := float64(NewAppData.GetCopysec())
	barCopy.Value = timeSecond
	barCopy.Min = 0.0
	barCopy.Max = timeSecond
	barCopy.Show()

	/* ticker := time.NewTicker(time.Second) */
	NewAppData.SetTicker(time.NewTicker(time.Second))
	for range NewAppData.GetTicker().C {
		timeSecond--
		fmt.Println("Left ", timeSecond)
		barCopy.SetValue(timeSecond)
		if timeSecond == 0.0 {
			barCopy.Hide()
			NewAppData.GetTicker().Stop()
		}
	}
}

func showPass(NewAppData *src.AppData, copyBtnLogin *widget.Button, copyBtnPass *widget.Button, login, pass string) {
	if copyBtnLogin.Text == cons.BTN_LABEL_COPY_LOGIN {
		openPass, _ := enigma.StartCrypt(pass, NewAppData.GetEntryCode().Text)

		copyBtnPass.SetText(openPass)
		copyBtnPass.Refresh()
		openLogin, _ := enigma.StartCrypt(login, NewAppData.GetEntryCode().Text)
		copyBtnLogin.SetText(openLogin)
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
	dialog.ShowConfirm("DELETE?", "REALY?", func(b bool) {
		if b {
			NewAppData.CellList = append(NewAppData.CellList[:id], NewAppData.CellList[id+1:]...)
			/* NewAppData.SetControlLen(len(NewAppData.CellList)) */
			SaveFile(NewAppData)
		}
	}, NewAppData.GetWindow())
}

func editCellDialog(NewAppData *src.AppData, id int) {
	if NewAppData.GetEntryCode().Text == "" {
		dialog.ShowInformation("Opps", "Please enter key-word", NewAppData.GetWindow())
		return
	} else {
		var newData [3]widget.Entry
		newData[0].PlaceHolder = "New Label"
		newData[1].PlaceHolder = "New Login"
		newData[2].PlaceHolder = "New Password"
		forms := container.NewVBox(&newData[0], &newData[1], &newData[2])
		dialog.ShowCustomConfirm("Edit", "Accept", "Exit", forms, func(b bool) {
			if b {
				editCell(id, newData, NewAppData)
			}
		}, NewAppData.GetWindow())
	}
}

func editCell(id int, newData [3]widget.Entry, NewAppData *src.AppData) {
	if newData[0].Text != "" {
		NewAppData.CellList[id].Label = newData[0].Text
	}
	if newData[1].Text != "" {
		s, b := enigma.StartCrypt(newData[1].Text, NewAppData.GetEntryCode().Text)
		if !b {
			return
		}
		NewAppData.CellList[id].Login = s
	}
	if newData[2].Text != "" {
		s, b := enigma.StartCrypt(newData[2].Text, NewAppData.GetEntryCode().Text)
		if !b {
			return
		}
		NewAppData.CellList[id].Pass = s
	}
	SaveFile(NewAppData)
}

func popUpMenu(NewAppData *src.AppData) *widget.PopUpMenu {
	popMenu := fyne.NewMenu("123", fyne.NewMenuItem("321", func() {}))
	pop := widget.NewPopUpMenu(popMenu, NewAppData.GetCanvas())
	return pop
}
