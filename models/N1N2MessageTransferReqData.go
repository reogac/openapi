package models
type N1N2MessageTransferReqData struct {
	 N1MessageContainer	*N1MessageContainer	`json:"n1MessageContainer,omitempty"`
	 SkipInd	*bool	`json:"skipInd,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 LastMsgIndication	*bool	`json:"lastMsgIndication,omitempty"`
	 Ppi	*int	`json:"ppi,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 ExtBufSupport	*bool	`json:"extBufSupport,omitempty"`
	 SmfReallocationInd	*bool	`json:"smfReallocationInd,omitempty"`
	 AreaOfValidity	*AreaOfValidity	`json:"areaOfValidity,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 OldGuami	*Guami	`json:"oldGuami,omitempty"`
	 N2InfoContainer	*N2InfoContainer	`json:"n2InfoContainer,omitempty"`
	 MtData	*RefToBinaryData	`json:"mtData,omitempty"`
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 LcsCorrelationId	string	`json:"lcsCorrelationId,omitempty"`
	 N1n2FailureTxfNotifURI	string	`json:"n1n2FailureTxfNotifURI,omitempty"`
	 TargetAccess	string	`json:"targetAccess,omitempty"`
}
