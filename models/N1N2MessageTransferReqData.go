package models

type N1N2MessageTransferReqData struct {
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
}
