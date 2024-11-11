package models

type N1N2MessageTransferReqData struct {
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	TargetAccess           string              `json:"targetAccess,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Fiveqi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
}
