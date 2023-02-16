package menu

import (
	"PassManager/confile"
	"PassManager/cons"
	"PassManager/menu/upd"
	/* "PassManager/src" */
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type InfaceApp interface {
	GetWindow() fyne.Window
	SetCopysec(int)
}

func GetMenu(iface InfaceApp) *fyne.MainMenu {
	menuBtnAbout := fyne.NewMenuItem(cons.MENU_BTN_ABOUT, func() { showVersionDalog(iface.GetWindow()) })
	menuBtnLargecopy := fyne.NewMenuItem(cons.MENU_BTN_LARGECOPY, nil)
	menuBtnLargecopy.ChildMenu = fyne.NewMenu("SubMenu",
		fyne.NewMenuItem(cons.SUBMENU_ONE, func() { setDurationCopy(iface, 5) }),
		fyne.NewMenuItem(cons.SUBMENU_TWO, func() { setDurationCopy(iface, 10) }),
		fyne.NewMenuItem(cons.SUBMENU_THREE, func() { setDurationCopy(iface, 15) }),
	)
	menuBtnNewBase := fyne.NewMenuItem(cons.MENU_BTN_NEWBASE, nil)

	menu := fyne.NewMenu("Menu", menuBtnNewBase, menuBtnLargecopy, menuBtnAbout)
	mainMenu := fyne.NewMainMenu(menu)
	return mainMenu
}

func showVersionDalog(w fyne.Window) {
	vers := confile.GetFromIni("data", "version")
	checkVersion := upd.ChekVersion()
	if vers == checkVersion {
		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", widget.NewLabel("Используется актуальная версия "+vers), w)
	} else {
		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", widget.NewLabel("Данная версия устарела. \n Актуальная версия: "+checkVersion), w)
	}
}

func setDurationCopy(iface InfaceApp, i int) {
	iface.SetCopysec(i)
	confile.SaveToIni("data", "duration", strconv.Itoa(i))
}
