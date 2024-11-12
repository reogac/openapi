package models

type AreaScope struct {
	TacInfoPerPlmn  map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
	EutraCellIdList []string           `json:"eutraCellIdList,omitempty"`
	NrCellIdList    []string           `json:"nrCellIdList,omitempty"`
	TacList         []string           `json:"tacList,omitempty"`
}
