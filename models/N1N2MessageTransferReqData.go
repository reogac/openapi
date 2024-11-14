package models
type N1N2MessageTransferReqData struct {
	 ExtBufSupport	*bool	`json:"extBufSupport,omitempty"`
	 TargetAccess	AccessType	`json:"targetAccess,omitempty"`
	 NfId	string	`json:"nfId,omitempty"`
	 SkipInd	*bool	`json:"skipInd,omitempty"`
	 LcsCorrelationId	string	`json:"lcsCorrelationId,omitempty"`
	 Ppi	*int	`json:"ppi,omitempty"`
	 MaAcceptedInd	*bool	`json:"maAcceptedInd,omitempty"`
	 N2InfoContainer	*N2InfoContainer	`json:"n2InfoContainer,omitempty"`
	 MtData	*RefToBinaryData	`json:"mtData,omitempty"`
	 FiveQi	*int	`json:"5qi,omitempty"`
	 N1n2FailureTxfNotifURI	string	`json:"n1n2FailureTxfNotifURI,omitempty"`
	 SmfReallocationInd	*bool	`json:"smfReallocationInd,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 OldGuami	*Guami	`json:"oldGuami,omitempty"`
	 LastMsgIndication	*bool	`json:"lastMsgIndication,omitempty"`
	 Arp	*Arp	`json:"arp,omitempty"`
	 AreaOfValidity	*AreaOfValidity	`json:"areaOfValidity,omitempty"`
	 N1MessageContainer	*N1MessageContainer	`json:"n1MessageContainer,omitempty"`
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
}
