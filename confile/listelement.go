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

	copyBtnPass := createBtnWithIcon(NewAppData, pass, cons.BTN_LABEL_COPY_PASS)
	copyBtnLogin := createBtnWithIcon(NewAppData, login, cons.BTN_LABEL_COPY_LOGIN)

	contManageCell := container.NewGridWithColumns(3,
		elem.NewButton("Edit", nil),
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
		container.NewCenter(container.New(
			layout.NewMaxLayout(),
			nameLabel,
			canvas.NewRectangle(color.RGBA{17, 0, 123, 1}),
		)),
		copyBtnLogin,
		copyBtnPass,
	)

	listElementContainer := container.NewVBox(line, contManageCell, contNameLogPass)

	return listElementContainer
}

func createBtnWithIcon(NewAppData *src.AppData, data, name string) *widget.Button {
	/* pop := widget.NewPopUp(widget.NewLabel("123"), NewAppData.GetCanvas()) */
	txtBoundPass := binding.NewString()
	txtBoundPass.Set(data)
	copyBtn := widget.NewButtonWithIcon(name, theme.ContentCopyIcon(), func() {
		if content, err := txtBoundPass.Get(); err == nil { // content - строка, которую надо расшифровать перед копированием или показом
			toCopy, errEnc := enigma.StartCrypt(content, NewAppData.GetEntryCode().Text)
			if !errEnc {
				dialog.ShowCustom("Error", "OK", widget.NewLabel(toCopy), NewAppData.GetWindow())
				return
			}
			NewAppData.GetWindow().Clipboard().SetContent(toCopy)
			/* pop.Show() */
			<-time.After(10 * time.Second)
			fmt.Println("Timer is gone")
			NewAppData.GetWindow().Clipboard().SetContent("")
		}

	})
	return copyBtn
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
	NewAppData.CellList = append(NewAppData.CellList[:id], NewAppData.CellList[id+1:]...)
	/* NewAppData.SetControlLen(len(NewAppData.CellList)) */
	SaveFile(NewAppData)
}
