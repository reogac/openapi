package models

type N1N2MessageTransferReqData struct {
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
}
