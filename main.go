package main

import (
	"encoding/json"
	"fmt"
	"io"

	"PassManager/cell"
	"PassManager/cons"
	"PassManager/elem"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	/* "fyne.io/fyne/v2/canvas" */
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	/* "fyne.io/fyne/v2/widget" */)

/* type CellData struct {
	Label string
	Login string
	Pass  string
} */

func main() {

	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME)
	canvas := mainWindow.Canvas()

	NewAppData := src.NewAppData(App, mainWindow, canvas)

	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	canvas.SetContent(container.NewCenter(createMangerBtns(NewAppData)))
	mainWindow.Show()
	App.Run()
}

func createMangerBtns(NewAppData *src.AppData) *fyne.Container {
	NewAppData.GetEntryCode().PlaceHolder = "Enter KeyCode"
	containerAddandKey := container.NewGridWithColumns(2, elem.NewButton(cons.BTN_LABEL_CREATE_NEW_CELL, func() {
		createNewCellList(NewAppData)
	}), NewAppData.GetEntryCode())
	containerOpenSaveBtn := container.NewGridWithColumns(2, elem.NewButton(cons.BTN_LABEL_OPEN, func() {
		openFile(NewAppData)
	}), elem.NewButton(cons.BTN_LABEL_SAVE, func() {
		saveFile(NewAppData)
	}))
	containerManager := container.NewGridWithRows(2, containerAddandKey, containerOpenSaveBtn)
	return containerManager
}

func createList(NewAppData *src.AppData) *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(NewAppData.CellList); i++ {
		containerListElement := elem.CreateListElement(NewAppData.CellList[i].Label, NewAppData.CellList[i].Login, NewAppData.CellList[i].Pass, NewAppData.GetWindow(), NewAppData.GetEntryCode().Text)
		listContainer.Add(containerListElement)
	}
	return listContainer
}

func createNewCellList(NewAppData *src.AppData) {
	newCell := cell.CreateNewCell()

	sendBtn := elem.NewButton("Save Data", func() { setDataFromDialogCell(newCell, NewAppData) })

	dialogContainer := container.NewVBox(newCell.GetLabel(), newCell.GetLogin(), newCell.GetPass(), sendBtn)

	dialog.ShowCustom(cons.DIALOG_CREATE_CELL_NAME, "Send", dialogContainer, NewAppData.GetWindow())
}

func setDataFromDialogCell(newCell *cell.Cell, NewAppData *src.AppData) {
	newCellData := src.NewCellData()

	newCellData.Label = newCell.GetLabel().Text
	newCellData.Login = newCell.GetLogin().Text
	newCellData.Pass = newCell.GetPass().Text

	NewAppData.CellList = append(NewAppData.CellList, *newCellData)

	NewAppData.GetCanvas().SetContent(container.NewVSplit(createMangerBtns(NewAppData), createList(NewAppData)))

	fmt.Println(NewAppData.CellList)
}

func saveFile(NewAppData *src.AppData) {
	code, err := json.Marshal(NewAppData.CellList)
	if err != nil {
		fmt.Println("Error", err)
	}

	dialog.ShowFileSave(
		func(uc fyne.URIWriteCloser, err error) {
			if uc != nil {
				io.WriteString(uc, string(code))
				NewAppData.GetCanvas().SetContent(container.NewVSplit(createMangerBtns(NewAppData), createList(NewAppData)))
			} else {
				return
			}
		}, NewAppData.GetWindow(),
	)
}

func openFile(NewAppData *src.AppData) {
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewAppData.CellList)
				if err != nil {
					panic(err)
				}

				NewAppData.GetCanvas().SetContent(container.NewVSplit(createMangerBtns(NewAppData), createList(NewAppData)))

			} else {
				return
			}
		}, NewAppData.GetWindow(),
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
