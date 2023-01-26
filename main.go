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
	/* "fyne.io/fyne/v2/canvas" */
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type CellData struct {
	Label string
	Login string
	Pass  string
}

var EntryCode widget.Entry

var NewCellList = []CellData{}
var Appp fyne.App

func main() {

	App := app.New()
	Appp = App
	mainWindow := App.NewWindow(cons.WINDOW_NAME)
	canvas := mainWindow.Canvas()

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	canvas.SetContent(container.NewVBox(createMangerBtns(mainWindow, canvas)))
	mainWindow.Show()
	App.Run()
}

/* func сreateWindowContent(mainWindow fyne.Window, canvas fyne.Canvas) *fyne.Container {
	return container.NewVBox(createMangerBtns(mainWindow, canvas))
} */

func createMangerBtns(mainWindow fyne.Window, canvas fyne.Canvas) *fyne.Container {
	containerAddandKey := container.NewGridWithColumns(2, elem.NewButton(cons.BTN_LABEL_CREATE_NEW_CELL, func() {
		createNewCellList(mainWindow)
	}), &EntryCode)
	containerOpenSaveBtn := container.NewGridWithColumns(2, elem.NewButton("Open", func() {
		openFile(mainWindow, canvas)
	}), elem.NewButton("Save", func() {
		saveFile(NewCellList, mainWindow)
	}))
	containerManager := container.NewGridWithRows(2, containerAddandKey, containerOpenSaveBtn)
	return containerManager
}

func createList(w *fyne.Window) *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(NewCellList); i++ {
		containerListElement := elem.CreateListElement(NewCellList[i].Label, NewCellList[i].Login, NewCellList[i].Pass, *w, EntryCode.Text)
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

func openFile(w fyne.Window, canvas fyne.Canvas) {
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewCellList)
				if err != nil {
					panic(err)
				}

				newContent := container.NewVBox(createMangerBtns(w, canvas), createList(&w))
				canvas.SetContent(newContent)

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
