package models
type Exception struct {
	 ExcepId	ExceptionId	`json:"excepId"`
	 ExcepLevel	*int	`json:"excepLevel,omitempty"`
	 ExcepTrend	ExceptionTrend	`json:"excepTrend,omitempty"`
}
