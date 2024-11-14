package models
type DnPerfInfo struct {
	 Dnn	string	`json:"dnn,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 DnPerf	[]DnPerf	`json:"dnPerf"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 AppId	string	`json:"appId,omitempty"`
}
