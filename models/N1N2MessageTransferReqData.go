package models

type N1N2MessageTransferReqData struct {
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
}
