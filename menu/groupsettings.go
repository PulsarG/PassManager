// GROUP SETTING AND MANAGE

package menu

import (
	"PassManager/confile"
	"PassManager/cons"
	"PassManager/src"
	// "fmt"

	// "PassManager/cons"
	"PassManager/elem"
	// "PassManager/menu/upd"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func createMenuGroupSettings(iface confile.InfaceApp) *fyne.MenuItem {

	menuBtnGroupSettings := fyne.NewMenuItem(cons.MENU_BTN_GROUP_SETTINGS, func() {
		var selectedGroup string
		selGroup := widget.NewSelect(
			getGroupList(iface.GetCellList()),
			func(s string) {
				selectedGroup = s
			},
		)

		newNameGr := widget.NewEntry()
		newNameGr.PlaceHolder = cons.DIALOG_GROUP_SET_ENTRY_NEWNAME_PLACEHOLDER

		btnGrDelete := elem.NewButton("Del", func() {
			deleteGroup(iface, selectedGroup)
		})
		btnGrRename := elem.NewButton("Rename", func() {
			renameGroup(iface, selectedGroup, newNameGr.Text)
		})

		groupSetCont := container.NewVBox(selGroup, newNameGr, btnGrDelete, btnGrRename)

		dialog.ShowCustom(cons.MENU_BTN_GROUP_SETTINGS, "Exit", groupSetCont, iface.GetWindow())
	})

	return menuBtnGroupSettings
}

func deleteGroup(iface confile.InfaceApp, nameGr string) {
	dialog.ShowConfirm(cons.DIALOG_DELETE_NAME, cons.DIALOG_DELETE_CONFIRM, func(b bool) {
		if b { // if outer
			val := iface.GetCellList()[nameGr]
			if len(val) == 0 { // if inner
				delete(iface.GetCellList(), nameGr)
				iface.GetCanvas().SetContent(container.NewHSplit(confile.CreateMangerBtns(iface), confile.CreateList(iface)))
				confile.SaveFile(iface)
			} else {
				iface.GetInfoDialog().ShowInfo("Group not empty")
			} // end if inner
		} // end if outer
	}, iface.GetWindow())
}

func renameGroup(iface confile.InfaceApp, nameGr, newName string) {
	var arrCell []src.CellData
	if newName != "" {
		arrCell = iface.GetCellList()[nameGr]
		delete(iface.GetCellList(), nameGr)
		iface.GetCellList()[newName] = arrCell
		confile.SaveFile(iface)
	} // end if
}

func getGroupList(m map[string][]src.CellData) []string {
	var groupp []string
	for gr, _ := range m {
		if m != nil {
			groupp = append(groupp, gr)
		} // end if
	} // end for
	return groupp
}
