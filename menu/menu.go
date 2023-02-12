package menu

import (
	"PassManager/cons"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GetMenu(NewAppData *src.AppData) *fyne.MainMenu {
	menuBtnAbout := fyne.NewMenuItem(cons.MENU_BTN_ABOUT, func() { showVersionDalog(NewAppData.GetWindow()) })
	menuBtnLargecopy := fyne.NewMenuItem(cons.MENU_BTN_LARGECOPY, nil)
	menuBtnLargecopy.ChildMenu = fyne.NewMenu("SubMenu",
		fyne.NewMenuItem(cons.SUBMENU_ONE, func() { setDurationCopy(NewAppData, 5.0) }),
		fyne.NewMenuItem(cons.SUBMENU_TWO, func() { setDurationCopy(NewAppData, 10.0) }),
		fyne.NewMenuItem(cons.SUBMENU_THREE, func() { setDurationCopy(NewAppData, 15.0) }),
	)

	menu := fyne.NewMenu("Menu", menuBtnLargecopy, menuBtnAbout)
	mainMenu := fyne.NewMainMenu(menu)
	return mainMenu
}

func showVersionDalog(w fyne.Window) {
	dialog.ShowCustom("Check Version", "Cancel", widget.NewLabel("Используется актуальная версия"), w)
}

func setDurationCopy(NewAppData *src.AppData, f float64) {
	NewAppData.SetCopysec(f)
}
