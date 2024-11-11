type AlternativeQosProfile struct {
	 Index	int	`json:"index"`
	 GuaFbrDl	string	`json:"guaFbrDl,omitempty"`
	 GuaFbrUl	string	`json:"guaFbrUl,omitempty"`
	 PacketDelayBudget	*int	`json:"packetDelayBudget,omitempty"`
	 PacketErrRate	string	`json:"packetErrRate,omitempty"`
}
