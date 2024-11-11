type VplmnQos struct {
	 SessionAmbr	*Ambr	`json:"sessionAmbr,omitempty"`
	 MaxFbrDl	string	`json:"maxFbrDl,omitempty"`
	 MaxFbrUl	string	`json:"maxFbrUl,omitempty"`
	 GuaFbrDl	string	`json:"guaFbrDl,omitempty"`
	 GuaFbrUl	string	`json:"guaFbrUl,omitempty"`
	 5qi	*int	`json:"5qi,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
}
