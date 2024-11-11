package models

type Exception struct {
	ExcepId    string `json:"excepId"`
	ExcepLevel *int   `json:"excepLevel,omitempty"`
	ExcepTrend string `json:"excepTrend,omitempty"`
}
