// ONE CELL -NAME-LOG-PASS CLASS

package src

type CellData struct {
	Label string
	Login string
	Pass  string
}

func NewCellData() *CellData {
	return &CellData{}
}
