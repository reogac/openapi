package models

type N1N2MessageTransferReqData struct {
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
}
