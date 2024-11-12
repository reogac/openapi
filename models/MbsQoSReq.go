package models

type MbsQoSReq struct {
	MaxBitRate  string `json:"maxBitRate,omitempty"`
	AverWindow  *int   `json:"averWindow,omitempty"`
	ReqMbsArp   *Arp   `json:"reqMbsArp,omitempty"`
	FiveQi      int    `json:"5qi"`
	GuarBitRate string `json:"guarBitRate,omitempty"`
}
