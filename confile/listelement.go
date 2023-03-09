package confile

import (
	"PassManager/cons"
	// "crypto/sha256"
	"fmt"
	"image/color"
	"time"

	"PassManager/elem"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma"
)

func createListElement(groupp string, id int, label, login, pass string, iface InfaceApp) *fyne.Container {
	copyBtnPass := createBtnWithIcon(iface, pass, cons.BTN_LABEL_COPY_PASS)
	copyBtnLogin := createBtnWithIcon(iface, login, cons.BTN_LABEL_COPY_LOGIN)

	contManageCell := container.NewHBox(
		createManageBtn(cons.BTN_LABEL_EDIT, func() {
			editCellDialog(iface, id, groupp)
		}),

		createManageBtn(cons.BTN_LABEL_DELETE, func() {
			deleteCell(id, iface, groupp)

			// !!! Test hach <
			/* h := sha256.New()
			h.Write([]byte(iface.GetEntryCode().Text))
			hashBytes := h.Sum(nil)
			hashStr := fmt.Sprintf("%x", hashBytes)
			fmt.Println(hashStr) */
			// !!! >
		}),

		createManageBtn(cons.BTN_LABEL_SHOW_LOGPASS, func() {
			showPass(iface, copyBtnLogin, copyBtnPass, login, pass)
		}),
	)

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 1

	nameLabel := widget.NewLabel(label)
	nameLabel.TextStyle.Bold = true
	nameLabel.TextStyle.Italic = true
	contWithName := container.NewCenter(nameLabel)

	contNameLogPass := container.NewGridWithColumns(4,
		contWithName,
		copyBtnLogin,
		copyBtnPass,
		contManageCell,
	)

	listElementContainer := container.NewVBox(line, contNameLogPass)

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

func createManageBtn(label string, f func()) *fyne.Container {
	btn := elem.NewButton(label, f)
	btn2 := container.New(
		layout.NewMaxLayout(),
		btn,
	)
	container := container.NewWithoutLayout(btn2)
	btn.Resize(fyne.NewSize(5, 5))
	return container
}

func createBtnWithIcon(iface InfaceApp, data, name string) *widget.Button {
	txtBoundPass := binding.NewString()
	txtBoundPass.Set(data)
	copyBtn := widget.NewButtonWithIcon(name, theme.ContentCopyIcon(), func() {
		if iface.GetTicker() != nil {
			iface.GetTicker().Stop()
		}
		go copyAndBarr(txtBoundPass, iface)
	})
	return copyBtn
}

func copyAndBarr(txtBoundPass binding.String, iface InfaceApp) {
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
	progressBarLine(iface)
}

func progressBarLine(iface InfaceApp) {
	timeSecond := float64(iface.GetCopysec())
	iface.GetMainBar().Value = timeSecond
	iface.GetMainBar().Min = 0.0
	iface.GetMainBar().Max = timeSecond
	iface.GetMainBar().Show()

	iface.SetTicker(time.NewTicker(time.Second))
	for range iface.GetTicker().C {
		iface.GetMainBar().SetValue(timeSecond)
		timeSecond--
		iface.GetMainBar().SetValue(timeSecond) // With additional output, the progress bar doesn't end too abruptly.
		if timeSecond <= 0.0 {
			iface.GetMainBar().Hide()
			iface.GetTicker().Stop()
			iface.GetWindow().Clipboard().SetContent("")
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

func CreateList(iface InfaceApp) *container.Scroll {

	acc := widget.NewAccordion()
	for gr, _ := range iface.GetCellList() {
		acc.Append(widget.NewAccordionItem(gr, createOneGroupp(iface, gr)))
	}
	return container.NewVScroll(acc)
}

func createOneGroupp(iface InfaceApp, gr string) *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(iface.GetCellList()[gr]); i++ {
		containerListElement := createListElement(gr, i, iface.GetCellList()[gr][i].Label, iface.GetCellList()[gr][i].Login, iface.GetCellList()[gr][i].Pass, iface)
		listContainer.Add(containerListElement)
	}
	return listContainer
}

func deleteCell(id int, iface InfaceApp, gr string) {
	dialog.ShowConfirm(cons.DIALOG_DELETE_NAME, cons.DIALOG_DELETE_CONFIRM, func(b bool) {
		if b {
			iface.SetDeleteCell(id, gr)
			SaveFile(iface)
		}
	}, iface.GetWindow())
}

func editCellDialog(iface InfaceApp, id int, gr string) {
	if iface.GetEntryCode().Text == "" {
		iface.GetInfoDialog().ShowInfo(cons.DIALOG_MESSAGE_NO_KEY)
		return
	} else {
		var newData [3]widget.Entry
		newData[0].PlaceHolder = "New Label"
		newData[1].PlaceHolder = "New Login"
		newData[2].PlaceHolder = "New Password"

		var groupp []string
		for gr, _ := range iface.GetCellList() {
			if iface.GetCellList() != nil {
				groupp = append(groupp, gr)
			}
		}
		selGroupp := widget.NewSelectEntry(
			groupp,
		)
		selGroupp.PlaceHolder = "select a group or enter a new one"

		// Заполнение формы существующим
		newData[0].SetText(iface.GetCellList()[gr][id].Label)
		logV, _ := enigma.StartCrypt(iface.GetCellList()[gr][id].Login, iface.GetEntryCode().Text)
		newData[1].SetText(logV)
		passV, _ := enigma.StartCrypt(iface.GetCellList()[gr][id].Pass, iface.GetEntryCode().Text)
		newData[2].SetText(passV)

		forms := container.NewVBox(&newData[0], &newData[1], &newData[2], selGroupp)
		dialog.ShowConfirm("Attention",
			cons.DIALOG_ATTENTION_EDIT_CELL_INFO,
			func(b bool) {
				if b {
					dialog.ShowCustomConfirm("Edit", "Accept", "Exit", forms, func(b bool) {
						if b {
							editCell(id, newData, iface, gr, selGroupp.Text)
						}
					}, iface.GetWindow())
				}
			}, iface.GetWindow())
	}
}

func editCell(id int, newData [3]widget.Entry, iface InfaceApp, gr, newGr string) {
	if newData[0].Text != "" {
		iface.GetCellList()[gr][id].Label = newData[0].Text
	}
	if newData[1].Text != "" {
		s, b := enigma.StartCrypt(newData[1].Text, iface.GetEntryCode().Text)
		if !b {
			return
		}
		iface.GetCellList()[gr][id].Login = s
	}
	if newData[2].Text != "" {
		s, b := enigma.StartCrypt(newData[2].Text, iface.GetEntryCode().Text)
		if !b {
			return
		}
		iface.GetCellList()[gr][id].Pass = s
	}

	if newGr != "" {
		cell := iface.GetCellList()[gr][id]
		iface.SetDeleteCell(id, gr)

		iface.SetCellListAppend(cell, newGr)
	}

	SaveFile(iface)
}

func popUpMenu(iface InfaceApp) *widget.PopUpMenu {
	popMenu := fyne.NewMenu("123", fyne.NewMenuItem("321", func() {}))
	pop := widget.NewPopUpMenu(popMenu, iface.GetCanvas())
	return pop
}
