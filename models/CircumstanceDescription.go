package models
type CircumstanceDescription struct {
	 LocArea	*NetworkAreaInfo	`json:"locArea,omitempty"`
	 Vol	*int64	`json:"vol,omitempty"`
	 Freq	*float64	`json:"freq,omitempty"`
	 Tm	string	`json:"tm,omitempty"`
}
