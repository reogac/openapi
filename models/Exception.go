package models

type Exception struct {
	ExcepTrend ExceptionTrend `json:"excepTrend,omitempty"`
	ExcepId    ExceptionId    `json:"excepId"`
	ExcepLevel *int           `json:"excepLevel,omitempty"`
}
