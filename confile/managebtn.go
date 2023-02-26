package confile

import (
	"encoding/json"
	"image/color"

	"PassManager/cons"
	"PassManager/elem"
	"PassManager/passgen"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/Enigma"
)

func CreateMangerBtns(iface InfaceApp) *fyne.Container {
	iface.GetEntryCode().PlaceHolder = cons.ENTER_KEY_PLACEHOLDER

	btnCreateCell := createColorBtn(cons.BTN_LABEL_CREATE_NEW_CELL, iface, func() { createNewCellList(iface) })

	btnOpen := createColorBtn(cons.BTN_LABEL_OPEN, iface, func() { OpenFile(iface) })

	btnOpenCustomRotor := elem.NewButton(cons.BTN_LABEL_OPEN_ROTOR, func() {
		GetRotorFromFile(iface)
	})

	btnCreateCustomRotor := createColorBtn(cons.BTN_LABEL_CREATE_CUSTOM_ROTOR, iface, func() {
		createSaveNewRotor(iface)
	})

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 1
	line2 := canvas.NewLine(color.Black)
	line.StrokeWidth = 1
	line3 := canvas.NewLine(color.Black)
	line.StrokeWidth = 1

	containerVbtns := container.NewVBox(iface.GetEntryCode(), line, btnCreateCell, line2, btnOpen, line3, btnOpenCustomRotor, btnCreateCustomRotor)

	bord := container.NewBorder(containerVbtns, iface.GetMainBar(), nil, nil)

	return bord
}

func createSaveNewRotor(iface InfaceApp) {
	rotor, errRotor := enigma.NewRotor()
	if !errRotor {
		dialog.ShowInformation("Error", "Opps, try again", iface.GetWindow())
	}
	rotorData, err := json.Marshal(rotor)
	if err != nil {
		dialog.ShowInformation("Error", "Opps, try again", iface.GetWindow())
	}
	createNewRotorFile(iface, rotorData)
}

func createColorBtn(label string, iface InfaceApp, f func()) *fyne.Container {
	return container.New(
		layout.NewMaxLayout(),
		canvas.NewRectangle(color.NRGBA{R: 11, G: 78, B: 150, A: 1}),
		elem.NewButton(label, f),
	)
}

func createNewCellList(iface InfaceApp) {
	if iface.GetEntryCode().Text != "" {
		newCell := src.CreateNewCell()
		form := widget.NewForm(
			widget.NewFormItem(cons.FORM_LABEL_NAME, newCell.GetLabel()),
			widget.NewFormItem(cons.FORM_LABEL_LOGIN, newCell.GetLogin()),
			widget.NewFormItem(cons.FORM_LABEL_PASS, newCell.GetPass()),
		)

		var groupp []string
		for gr, _ := range iface.GetCellList() {
			if iface.GetCellList() != nil {
				groupp = append(groupp, gr)
			}
		}
		selGroupp := widget.NewSelectEntry(
			groupp,
		)

		comt := container.NewVBox(form, selGroupp, elem.NewButton("Random pass (20)", func() {
			newCell.GetPass().SetText(passgen.GetRandomPass())
		}))
		dialog.ShowCustomConfirm(cons.DIALOG_CREATE_CELL_NAME,
			"Add",
			"Close",
			comt, func(b bool) {
				if b {
					setDataFromDialogCell(newCell, iface, selGroupp.Text)
				} else {
					return
				}
			},
			iface.GetWindow())
	} else {
		dialog.ShowCustom("Oops", "Ok", widget.NewLabel("Please entry Key-Word"), iface.GetWindow())
	}
}

func setDataFromDialogCell(newCell *src.Cell, iface InfaceApp, groupp string) {
	newCellData := src.NewCellData()
	var err bool

	newCellData.Label = newCell.GetLabel().Text
	newCellData.Login, err = enigma.StartCrypt(newCell.GetLogin().Text, iface.GetEntryCode().Text)
	if !err {
		dialog.ShowCustom("Error", "OK", widget.NewLabel(newCellData.Login), iface.GetWindow())
		return
	}
	newCellData.Pass, err = enigma.StartCrypt(newCell.GetPass().Text, iface.GetEntryCode().Text)
	if !err {
		dialog.ShowCustom("Error", "OK", widget.NewLabel(newCellData.Pass), iface.GetWindow())
		return
	}

	iface.SetCellListAppend(*newCellData, groupp)

	iface.GetCanvas().SetContent(container.NewHSplit(CreateMangerBtns(iface), CreateList(iface)))

	SaveFile(iface)
}
