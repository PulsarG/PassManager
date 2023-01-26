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

type AppData struct {
	app         fyne.App
	newCellList []CellData
	entryCode   widget.Entry
	canvas      fyne.Canvas
	mainWindow  fyne.Window
}

func main() {

	App := app.New()
	mainWindow := App.NewWindow(cons.WINDOW_NAME)
	canvas := mainWindow.Canvas()

	NewAppData := AppData{app: App, mainWindow: mainWindow, canvas: canvas}
	
	mainWindow.Resize(fyne.NewSize(cons.WINDOW_MAIN_WEIGHT, cons.WINDOW_MAIN_HIGHT))

	canvas.SetContent(container.NewCenter(createMangerBtns(NewAppData)))
	mainWindow.Show()
	App.Run()
}

/* func сreateWindowContent(mainWindow fyne.Window, canvas fyne.Canvas) *fyne.Container {
	return container.NewVBox(createMangerBtns(mainWindow, canvas))
} */

func createMangerBtns(NewAppData AppData) *fyne.Container {
	NewAppData.entryCode.PlaceHolder = "Enter KeyCode"
	containerAddandKey := container.NewGridWithColumns(2, elem.NewButton(cons.BTN_LABEL_CREATE_NEW_CELL, func() {
		createNewCellList(NewAppData)
	}), &NewAppData.entryCode)
	containerOpenSaveBtn := container.NewGridWithColumns(2, elem.NewButton(cons.BTN_LABEL_OPEN, func() {
		openFile(NewAppData)
	}), elem.NewButton(cons.BTN_LABEL_SAVE, func() {
		/* saveFile(NewAppData) */
	}))
	containerManager := container.NewGridWithRows(2, containerAddandKey, containerOpenSaveBtn)
	return containerManager
}

func createList(NewAppData AppData) *fyne.Container {
	listContainer := container.NewVBox()
	for i := 0; i < len(NewAppData.newCellList); i++ {
		containerListElement := elem.CreateListElement(NewAppData.newCellList[i].Label, NewAppData.newCellList[i].Login, NewAppData.newCellList[i].Pass, *&NewAppData.mainWindow, NewAppData.entryCode.Text)
		listContainer.Add(containerListElement)
	}
	return listContainer
}

func createNewCellList(NewAppData AppData) {
	newCell := cell.CreateNewCell()

	sendBtn := elem.NewButton("Save Data", func() { setDataFromDialogCell(newCell, NewAppData) })

	dialogContainer := container.NewVBox(newCell.GetLabel(), newCell.GetLogin(), newCell.GetPass(), sendBtn)

	dialog.ShowCustom(cons.DIALOG_CREATE_CELL_NAME, "Send", dialogContainer, NewAppData.mainWindow)
}

func setDataFromDialogCell(newCell *cell.Cell, NewAppData AppData) {
	newCellData := CellData{}

	newCellData.Label = newCell.GetLabel().Text
	newCellData.Login = newCell.GetLogin().Text
	newCellData.Pass = newCell.GetPass().Text

	NewAppData.newCellList = append(NewAppData.newCellList, newCellData)

	fmt.Println(NewAppData.newCellList)
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

func openFile(NewAppData AppData) {
	dialog.ShowFileOpen(
		func(uc fyne.URIReadCloser, _ error) {
			if uc != nil {
				data, _ := io.ReadAll(uc)
				err := json.Unmarshal(data, &NewAppData.newCellList)
				if err != nil {
					panic(err)
				}

				NewAppData.canvas.SetContent(container.NewVSplit(createMangerBtns(NewAppData), createList(NewAppData)))

			} else {
				return
			}
		}, NewAppData.mainWindow,
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
