package models

type N1N2MessageTransferReqData struct {
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
}
