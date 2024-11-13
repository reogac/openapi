package models

type MbsQoSReq struct {
	ReqMbsArp   *Arp   `json:"reqMbsArp,omitempty"`
	FiveQi      int    `json:"5qi"`
	GuarBitRate string `json:"guarBitRate,omitempty"`
	MaxBitRate  string `json:"maxBitRate,omitempty"`
	AverWindow  *int   `json:"averWindow,omitempty"`
}
