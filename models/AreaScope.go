package models

type AreaScope struct {
	NrCellIdList    []string           `json:"nrCellIdList,omitempty"`
	TacList         []string           `json:"tacList,omitempty"`
	TacInfoPerPlmn  map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
	EutraCellIdList []string           `json:"eutraCellIdList,omitempty"`
}
