package models

type N1N2MessageTransferReqData struct {
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
}
