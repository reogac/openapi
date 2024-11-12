package models

type DispersionInfo struct {
	DisperType     DispersionType         `json:"disperType"`
	TsStart        string                 `json:"tsStart"`
	TsDuration     int                    `json:"tsDuration"`
	DisperCollects []DispersionCollection `json:"disperCollects"`
}
