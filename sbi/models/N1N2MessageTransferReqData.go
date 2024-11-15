/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	SkipInd                *bool               `json:"skipInd,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
}
