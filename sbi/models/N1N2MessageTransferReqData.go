/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
}
