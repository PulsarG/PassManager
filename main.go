package main

import (
	/* "fmt"
	"math/rand" */
	"encoding/json"
	"fmt"
	"io"

	/* "fyne.io/fyne/v2" */
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	/* "fyne.io/fyne/v2/layout" */
	"fyne.io/fyne/v2/widget"
	/* "fyne.io/fyne/v2/canvas" */ /* "github.com/PulsarG/Enigma_algorithm" */)

type Test struct {
	ID       int
	Username string
}

func main() {

	jsonStr := `{"id": 42, "username": "alex"}`

	data := []byte(jsonStr)

	t := &Test{}
	json.Unmarshal(data, t)
	fmt.Println(t.ID)

	textField := widget.NewEntry()

	App := app.New()
	w := App.NewWindow("qwe")

	w.Resize(fyne.NewSize(500, 500))

	t.ID = 50
	result, _ := json.Marshal(t)

	scroll := container.NewVBox(widget.NewCheck("qwe", nil), textField)

	manageConteiner := container.NewHBox(widget.NewButton("SSS", func() { saveFile(string(result), w) }), widget.NewButton("OOO", func() { openFile(textField, w) }))

	mainContainer := container.NewGridWithRows(2, scroll, manageConteiner)

	w.SetContent(mainContainer)
	w.Show()
	App.Run()
	/* Arr := make([]int, 10)

	for i := 0; i < len(Arr); i++ {

		readyIndex := make([]int, 10)

	AGAIN:
		j := rand.Intn(10)

		for l := 0; l < len(readyIndex); l++ {
			if readyIndex[l] == j {
				goto AGAIN
			} else {
				continue
			}
		}

		if i != j {
			Arr[i] = j
			Arr[j] = i
			readyIndex = append(readyIndex, j)
			continue
		} else {
			goto AGAIN
		}
	}
	for k := 0; k < len(Arr); k++ {
		fmt.Println("Индекс: ", k, "Значение: ", Arr[k])
	} */
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
	var d string
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, err error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)
				d = string(data)
				textField.SetText(d)
			} else {
				return
			}
		}, w,
	)
}
