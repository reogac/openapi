package models

type CongestionInfo struct {
	CongType     string           `json:"congType"`
	TimeIntev    TimeWindow       `json:"timeIntev"`
	Nsi          ThresholdLevel   `json:"nsi"`
	Confidence   *int             `json:"confidence,omitempty"`
	TopAppListUl []TopApplication `json:"topAppListUl,omitempty"`
	TopAppListDl []TopApplication `json:"topAppListDl,omitempty"`
}