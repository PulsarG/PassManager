package menu

import (
	"PassManager/confile"
	"PassManager/cons"
	"PassManager/menu/upd"
	"PassManager/src"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GetMenu(NewAppData *src.AppData) *fyne.MainMenu {
	menuBtnAbout := fyne.NewMenuItem(cons.MENU_BTN_ABOUT, func() { showVersionDalog(NewAppData.GetWindow()) })
	menuBtnLargecopy := fyne.NewMenuItem(cons.MENU_BTN_LARGECOPY, nil)
	menuBtnLargecopy.ChildMenu = fyne.NewMenu("SubMenu",
		fyne.NewMenuItem(cons.SUBMENU_ONE, func() { setDurationCopy(NewAppData, 5) }),
		fyne.NewMenuItem(cons.SUBMENU_TWO, func() { setDurationCopy(NewAppData, 10) }),
		fyne.NewMenuItem(cons.SUBMENU_THREE, func() { setDurationCopy(NewAppData, 15) }),
	)

	menu := fyne.NewMenu("Menu", menuBtnLargecopy, menuBtnAbout)
	mainMenu := fyne.NewMainMenu(menu)
	return mainMenu
}

func showVersionDalog(w fyne.Window) {
	vers := upd.GetVersion()
	checkVersion := upd.ChekVersion()
	if vers == checkVersion {
		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", widget.NewLabel("Используется актуальная версия "+vers), w)
	} else {
		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", widget.NewLabel("Данная версия устарела. \n Актуальная версия: "+checkVersion), w)
	}
}

func setDurationCopy(NewAppData *src.AppData, i int) {
	NewAppData.SetCopysec(i)
	confile.SaveCopysecIni(i)
}
