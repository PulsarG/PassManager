package main

import (
	"encoding/json"
	"fmt"
	"io"

	"PassManager/cell"
	"PassManager/cons"
	"PassManager/elem"
	/* 	"PassManager/src" */

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"

	/* "fyne.io/fyne/v2/canvas" */
	/* "fyne.io/fyne/v2/container" */
	"fyne.io/fyne/v2/dialog"
)

type CellData struct {
	Label string
	Login string
	Pass  string
}

var NewCellList = []CellData{}
var Appp fyne.App

func main() {
	
	App := app.New()
	Appp = App
	mainWindow := App.NewWindow(cons.WINDOW_NAME)

	/* windowContent := сreateWindowContent(mainWindow) */

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	mainWindow.SetContent(сreateWindowContent(mainWindow))
	mainWindow.Show()
	App.Run()
}

func сreateWindowContent(mainWindow fyne.Window) *fyne.Container {

	/* containerList := createList() */

	containerAddandKey := container.NewGridWithColumns(2, elem.NewButton(cons.BTN_LABEL_CREATE_NEW_CELL, func() {
		createNewCellList(mainWindow)
	}), widget.NewCheck("123", nil))
	containerOpenSaveBtn := container.NewGridWithColumns(2, elem.NewButton("Open", func() {
		openFile(mainWindow)
	}), elem.NewButton("Save", func() {
		saveFile(NewCellList, mainWindow)
	}))
	containerManager := container.NewGridWithRows(2, containerAddandKey, containerOpenSaveBtn)

	containerFull := container.NewCenter(containerManager)

	return containerFull
}

func createList() *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(NewCellList); i++ {
		containerListElement := elem.CreateListElement(NewCellList[i].Label, NewCellList[i].Login, NewCellList[i].Pass)
		listContainer.Add(containerListElement)
	}
	return listContainer
}

func createNewCellList(window fyne.Window) {
	newCell := cell.CreateNewCell()

	sendBtn := elem.NewButton("Save Data", func() { setDataFromDialogCell(newCell) })

	dialogContainer := container.NewVBox(newCell.GetLabel(), newCell.GetLogin(), newCell.GetPass(), sendBtn)

	dialog.ShowCustom(cons.DIALOG_CREATE_CELL_NAME, "Send", dialogContainer, window)
}

func setDataFromDialogCell(newCell *cell.Cell) {
	newCellData := CellData{}

	newCellData.Label = newCell.GetLabel().Text
	newCellData.Login = newCell.GetLogin().Text
	newCellData.Pass = newCell.GetPass().Text

	NewCellList = append(NewCellList, newCellData)

	fmt.Println(NewCellList)
}

func saveFile(NewCellList []CellData, w fyne.Window) {
	code, err := json.Marshal(NewCellList)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(NewCellList)
	fmt.Println(string(code))

	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				io.WriteString(uc, string(code))
			} else {
				return
			}
		}, w,
	)
}

func openFile(w fyne.Window) {

	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewCellList)
				if err != nil {
					panic(err)
				}

				/* textField.SetText(dd[0].Username) */

				for i := 0; i < len(NewCellList); i++ {
					fmt.Println("Unmarshal: ", NewCellList[i])
				}

				w2 := Appp.NewWindow("2134") // !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
				w2.SetContent(createList())
				w2.Show()

			} else {
				return
			}
		}, w,
	)
}

/* Старт приложения:

Проверка наличия существующего файла

Если файл есть - открытие файла - запрос на загрузку - выбор
	Парсинг данных из файла в список
		--- создание виджета списка
		--- сборка из Лейбла, Кнопок Логин и Пароль, кнопок Показа и Изменения
	Для добавления нового эелемента постоянная кнопка добавления
	с Диалоговым окном и и полями Лейбл, Логин, Пасс

Если файла нет - добавления нового эелемента постоянная кнопка добавления
с Диалоговым окном и и полями Лейбл, Логин, Пасс
	Внесение в БД
		Сохранение во внешнем файле */
