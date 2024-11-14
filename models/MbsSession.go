/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsSession struct {
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
}
