package models
type AreaScope struct {
	 EutraCellIdList	[]string	`json:"eutraCellIdList,omitempty"`
	 NrCellIdList	[]string	`json:"nrCellIdList,omitempty"`
	 TacList	[]string	`json:"tacList,omitempty"`
	 TacInfoPerPlmn	map[string]TacInfo	`json:"tacInfoPerPlmn,omitempty"`
}
