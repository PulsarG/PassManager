// CREATE MENU

package menu

import (
	"PassManager/confile"
	"PassManager/cons"
	// "PassManager/elem"
	"PassManager/menu/upd"
	"fmt"

	"net/url"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GetMenu(iface confile.InfaceApp) *fyne.MainMenu {
	menuBtnAbout := fyne.NewMenuItem(cons.MENU_BTN_ABOUT, func() { showVersionDalog(iface) })
	menuBtnLargecopy := fyne.NewMenuItem(cons.MENU_BTN_LARGECOPY, nil)
	menuBtnLargecopy.ChildMenu = fyne.NewMenu("SubMenu",
		fyne.NewMenuItem(cons.SUBMENU_ONE, func() { setDurationCopy(iface, 5) }),
		fyne.NewMenuItem(cons.SUBMENU_TWO, func() { setDurationCopy(iface, 10) }),
		fyne.NewMenuItem(cons.SUBMENU_THREE, func() { setDurationCopy(iface, 15) }),
	)
	menuBtnNewBase := fyne.NewMenuItem(cons.MENU_BTN_NEWBASE, func() {
		confile.SaveToIni("file", "path", "")
		iface.SetCellList(nil)
		confile.SaveFile(iface)
	})

	// ******
	subMenuSelectTrayHide := fyne.NewMenuItem(cons.SUBMENU_ONE_SET_TRAY, func() { confile.SaveToIni("data", "close", "false") })
	subMenuSelectTrayClose := fyne.NewMenuItem(cons.SUBMENU_ONE_SET_CLOSE, func() { confile.SaveToIni("data", "close", "true") })

	menuBtnSelectTray := fyne.NewMenuItem(cons.MENU_BTN_SELECT_TRAY_SYS, nil)

	menuBtnSelectTray.ChildMenu = fyne.NewMenu("SubMenu", subMenuSelectTrayHide, subMenuSelectTrayClose)

	menu := fyne.NewMenu("Menu", menuBtnNewBase, menuBtnLargecopy, createMenuGroupSettings(iface), menuBtnSelectTray, menuBtnAbout)

	mainMenu := fyne.NewMainMenu(menu)


	isSelected := confile.GetFromIni("data", "close")
	if isSelected == "true" {
		subMenuSelectTrayClose.Checked = true
		subMenuSelectTrayHide.Checked = false
		mainMenu.Refresh()
	} else if isSelected == "false" {
		subMenuSelectTrayClose.Checked = false
		subMenuSelectTrayHide.Checked = true
		mainMenu.Refresh()
	} else {
		subMenuSelectTrayClose.Checked = false
		subMenuSelectTrayHide.Checked = false
	}

	
	return mainMenu
}

func showVersionDalog(iface confile.InfaceApp) {
	vers := confile.GetFromIni("data", "version")
	checkVersion := upd.ChekVersion()
	if vers == checkVersion {
		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", widget.NewLabel(cons.MENU_UPDATE_ACTUAL+vers), iface.GetWindow())
	} else {
		url, errParse := url.Parse(cons.URL_GITHUB_LATEST_PAGE)
		if errParse != nil {
			fmt.Println("Parse fail")
		}
		container := container.NewVBox(
			widget.NewLabel(cons.MENU_UPDATE_OLD+checkVersion),
			widget.NewHyperlink(cons.MENU_OPEN_GITHUB_LINK, url),
			widget.NewButton("Update now", func() {
				dialog.ShowInformation("Update", upd.Update(), iface.GetWindow())
			}))

		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", container, iface.GetWindow())
	}
}

func setDurationCopy(iface confile.InfaceApp, i int) {
	iface.SetCopysec(i)
	confile.SaveToIni("data", "duration", strconv.Itoa(i))
}
