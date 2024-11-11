package models

type N1N2MessageTransferReqData struct {
	Ppi                    *int                `json:"ppi,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	TargetAccess           string              `json:"targetAccess,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
}
