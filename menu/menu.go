// CREATE MENU

package menu

import (
	// "PassManager/errs"
	"PassManager/errloger"
	"PassManager/confile"
	"PassManager/cons"
	// "PassManager/elem"
	"PassManager/menu/upd"
	// "fmt"

	"net/url"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/PulsarG/ConfigManager"
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
		inihandler.SaveToIni("file", "path", "")
		iface.SetCellList(nil)
		confile.SaveFile(iface)
	})

	// ******
	subMenuSelectTrayHide := fyne.NewMenuItem(cons.SUBMENU_ONE_SET_TRAY, func() { inihandler.SaveToIni("data", "close", "false") })
	subMenuSelectTrayClose := fyne.NewMenuItem(cons.SUBMENU_ONE_SET_CLOSE, func() { inihandler.SaveToIni("data", "close", "true") })

	menuBtnSelectTray := fyne.NewMenuItem(cons.MENU_BTN_SELECT_TRAY_SYS, nil)

	menuBtnSelectTray.ChildMenu = fyne.NewMenu("SubMenu", subMenuSelectTrayHide, subMenuSelectTrayClose)

	//*****
	menuBtnSelectLog := fyne.NewMenuItem("Send error to email?", func() { showLogSettingDalog(iface) })
	
	//*****

	menu := fyne.NewMenu("Menu", menuBtnNewBase, menuBtnLargecopy, createMenuGroupSettings(iface), menuBtnSelectTray, menuBtnSelectLog, menuBtnAbout)

	mainMenu := fyne.NewMainMenu(menu)

	isSelected := inihandler.GetFromIni("data", "close")
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
	vers := inihandler.GetFromIni("data", "version")
	checkVersion := upd.ChekVersion()
	if vers == checkVersion {
		dialog.ShowCustom(cons.MENU_BTN_ABOUT, "Cancel", widget.NewLabel(cons.MENU_UPDATE_ACTUAL+vers), iface.GetWindow())
	} else {
		url, errParse := url.Parse(cons.URL_GITHUB_LATEST_PAGE)
		if errParse != nil {
			errloger.ErrorLog(errParse)
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

func showLogSettingDalog(iface confile.InfaceApp) {
	dialog.ShowConfirm("Send errors to email?", "Отправлять информацию об ошибках\n разработчику?\n Будет отправлен только текст ошибки", nil, iface.GetWindow())
}

func setDurationCopy(iface confile.InfaceApp, i int) {
	iface.SetCopysec(i)
	inihandler.SaveToIni("data", "duration", strconv.Itoa(i))
}
