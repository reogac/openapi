package models
type AreaScope struct {
	 NrCellIdList	[]string	`json:"nrCellIdList,omitempty"`
	 TacList	[]string	`json:"tacList,omitempty"`
	 TacInfoPerPlmn	*tacInfoPerPlmn	`json:"tacInfoPerPlmn,omitempty"`
	 EutraCellIdList	[]string	`json:"eutraCellIdList,omitempty"`
}
