package models

type NcgiTai struct {
	CellList []Ncgi `json:"cellList"`
	Tai      Tai    `json:"tai"`
}
