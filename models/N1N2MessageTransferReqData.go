package models

type N1N2MessageTransferReqData struct {
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
}
