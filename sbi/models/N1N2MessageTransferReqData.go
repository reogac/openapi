/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
}
