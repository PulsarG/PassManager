package src

import (
	"PassManager/cons"

	"fyne.io/fyne/v2/widget"
)

type Cell struct {
	label widget.Entry
	login widget.Entry
	pass  widget.Entry
}

func CreateNewCell() *Cell {
	return &Cell{
	}
}

func (c *Cell) GetLabel() *widget.Entry {
	c.label.PlaceHolder = cons.DIALOG_CREATE_LABEL_PLACEHOLDER
	return &c.label
}

func (c *Cell) GetLogin() *widget.Entry {
	c.login.PlaceHolder = cons.DIALOG_CREATE_LOGIN_PLACEHOLDER
	return &c.login
}
func (c *Cell) GetPass() *widget.Entry {
	c.pass.PlaceHolder = cons.DIALOG_CREATE_PASS_PLACEHOLDER
	return &c.pass
}
