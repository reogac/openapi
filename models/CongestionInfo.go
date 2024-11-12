package models
type CongestionInfo struct {
	 Confidence	*int	`json:"confidence,omitempty"`
	 TopAppListUl	[]TopApplication	`json:"topAppListUl,omitempty"`
	 TopAppListDl	[]TopApplication	`json:"topAppListDl,omitempty"`
	 CongType	CongestionType	`json:"congType"`
	 TimeIntev	TimeWindow	`json:"timeIntev"`
	 Nsi	ThresholdLevel	`json:"nsi"`
}
