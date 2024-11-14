package models
type N1N2MessageTransferReqData struct {
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 SmfReallocationInd	*bool	`json:"smfReallocationInd,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 TargetAccess	AccessType	`json:"targetAccess,omitempty"`
	 N1MessageContainer	*N1MessageContainer	`json:"n1MessageContainer,omitempty"`
	 MtData	*RefToBinaryData	`json:"mtData,omitempty"`
	 N1n2FailureTxfNotifURI	string	`json:"n1n2FailureTxfNotifURI,omitempty"`
	 AreaOfValidity	*AreaOfValidity	`json:"areaOfValidity,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 ExtBufSupport	*bool	`json:"extBufSupport,omitempty"`
	 NfId	string	`json:"nfId,omitempty"`
	 LastMsgIndication	*bool	`json:"lastMsgIndication,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 OldGuami	*Guami	`json:"oldGuami,omitempty"`
	 SkipInd	*bool	`json:"skipInd,omitempty"`
	 Ppi	*int	`json:"ppi,omitempty"`
	 N2InfoContainer	*N2InfoContainer	`json:"n2InfoContainer,omitempty"`
	 LcsCorrelationId	string	`json:"lcsCorrelationId,omitempty"`
}
