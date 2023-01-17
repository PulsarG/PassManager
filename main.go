package main

import (
	"encoding/json"
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	/* "fyne.io/fyne/v2/layout" */
	"fyne.io/fyne/v2/widget"
	/* "fyne.io/fyne/v2/canvas" */)

type Test struct {
	ID       int
	Username string
}

type ArrTest struct {
	TestItems []Test
}

func main() {

	id := 2
	name := "Arve"

	test := Test{ID: id, Username: name}
	test2 := Test{ID: 2, Username: "Alex"}

	Ti := []Test{}
	Ti = append(Ti, test, test2)

	/* fmt.Println("Одиночный: ", test)
	fmt.Println("Слайс: ", Ti) */

	byteArray, _ := json.Marshal(Ti)

	fmt.Println(string(byteArray))

	textField := widget.NewEntry()

	App := app.New()
	w := App.NewWindow("qwe")

	w.Resize(fyne.NewSize(500, 500))

	scroll := container.NewVBox(widget.NewCheck("qwe", nil), textField)

	manageConteiner := container.NewHBox(widget.NewButton("SSS", func() { saveFile(string(byteArray), w) }), widget.NewButton("OOO", func() { openFile(textField, w) }))

	mainContainer := container.NewGridWithRows(2, scroll, manageConteiner)

	w.SetContent(mainContainer)
	w.Show()
	App.Run()
}

func saveFile(code string, w fyne.Window) {
	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				io.WriteString(uc, code)
			} else {
				return
			}
		}, w,
	)

}

func openFile(textField *widget.Entry, w fyne.Window) {
	/* var d string */
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, err error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)

				/* d = string(data) */

				dd := []Test{}
				err := json.Unmarshal(data, &dd)
				if err != nil {
					panic(err)
				}

				textField.SetText(dd[0].Username)

				for i := 0; i < len(dd); i++ {
					fmt.Println("Unmarshal: ", dd[i])
				}
			} else {
				return
			}
		}, w,
	)
}
