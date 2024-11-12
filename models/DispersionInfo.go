package models

type DispersionInfo struct {
	DisperCollects []DispersionCollection `json:"disperCollects"`
	DisperType     DispersionType         `json:"disperType"`
	TsStart        string                 `json:"tsStart"`
	TsDuration     int                    `json:"tsDuration"`
}
