package models
type MbsQoSReq struct {
	 FiveQi	int	`json:"5qi"`
	 GuarBitRate	string	`json:"guarBitRate,omitempty"`
	 MaxBitRate	string	`json:"maxBitRate,omitempty"`
	 AverWindow	*int	`json:"averWindow,omitempty"`
	 ReqMbsArp	*Arp	`json:"reqMbsArp,omitempty"`
}
