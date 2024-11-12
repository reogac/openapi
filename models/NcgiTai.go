package models

type NcgiTai struct {
	Tai      Tai    `json:"tai"`
	CellList []Ncgi `json:"cellList"`
}
