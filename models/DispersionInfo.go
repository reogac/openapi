package models
type DispersionInfo struct {
	 TsDuration	int	`json:"tsDuration"`
	 DisperCollects	[]DispersionCollection	`json:"disperCollects"`
	 DisperType	DispersionType	`json:"disperType"`
	 TsStart	string	`json:"tsStart"`
}
