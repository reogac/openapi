package models
type Exception struct {
	 ExcepLevel	*int	`json:"excepLevel,omitempty"`
	 ExcepTrend	ExceptionTrend	`json:"excepTrend,omitempty"`
	 ExcepId	ExceptionId	`json:"excepId"`
}
