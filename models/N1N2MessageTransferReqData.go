package models

type N1N2MessageTransferReqData struct {
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	TargetAccess           string              `json:"targetAccess,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
}
