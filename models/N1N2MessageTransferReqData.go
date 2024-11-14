/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MessageTransferReqData struct {
	SkipInd                *bool               `json:"skipInd,omitempty"`
	NfId                   string              `json:"nfId,omitempty"`
	Arp                    *Arp                `json:"arp,omitempty"`
	FiveQi                 *int                `json:"5qi,omitempty"`
	N1n2FailureTxfNotifURI string              `json:"n1n2FailureTxfNotifURI,omitempty"`
	OldGuami               *Guami              `json:"oldGuami,omitempty"`
	ExtBufSupport          *bool               `json:"extBufSupport,omitempty"`
	MtData                 *RefToBinaryData    `json:"mtData,omitempty"`
	PduSessionId           *int                `json:"pduSessionId,omitempty"`
	LcsCorrelationId       string              `json:"lcsCorrelationId,omitempty"`
	TargetAccess           AccessType          `json:"targetAccess,omitempty"`
	N1MessageContainer     *N1MessageContainer `json:"n1MessageContainer,omitempty"`
	N2InfoContainer        *N2InfoContainer    `json:"n2InfoContainer,omitempty"`
	MaAcceptedInd          *bool               `json:"maAcceptedInd,omitempty"`
	AreaOfValidity         *AreaOfValidity     `json:"areaOfValidity,omitempty"`
	SupportedFeatures      string              `json:"supportedFeatures,omitempty"`
	LastMsgIndication      *bool               `json:"lastMsgIndication,omitempty"`
	Ppi                    *int                `json:"ppi,omitempty"`
	SmfReallocationInd     *bool               `json:"smfReallocationInd,omitempty"`
}
