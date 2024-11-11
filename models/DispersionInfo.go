package models

type DispersionInfo struct {
	TsStart        string                 `json:"tsStart"`
	TsDuration     int                    `json:"tsDuration"`
	DisperCollects []DispersionCollection `json:"disperCollects"`
	DisperType     string                 `json:"disperType"`
}
