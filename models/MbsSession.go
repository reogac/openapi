package models

type MbsSession struct {
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
}
