package models

type N1N2MessageTransferReqData struct {
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
}
