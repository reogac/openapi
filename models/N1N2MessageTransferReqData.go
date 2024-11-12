package models

type N1N2MessageTransferReqData struct {
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
}
