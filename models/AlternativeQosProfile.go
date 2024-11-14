package models
type AlternativeQosProfile struct {
	 PacketDelayBudget	*int	`json:"packetDelayBudget,omitempty"`
	 PacketErrRate	string	`json:"packetErrRate,omitempty"`
	 Index	int	`json:"index"`
	 GuaFbrDl	string	`json:"guaFbrDl,omitempty"`
	 GuaFbrUl	string	`json:"guaFbrUl,omitempty"`
}
