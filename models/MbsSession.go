package models

type MbsSession struct {
	MbsFsaIdList      []string                 `json:"mbsFsaIdList,omitempty"`
	Tmgi              *Tmgi                    `json:"tmgi,omitempty"`
	ServiceType       MbsServiceType           `json:"serviceType"`
	LocationDependent *bool                    `json:"locationDependent,omitempty"`
	AreaSessionId     *int                     `json:"areaSessionId,omitempty"`
	ExtRedMbsServArea *ExternalMbsServiceArea  `json:"extRedMbsServArea,omitempty"`
	MbsServInfo       *MbsServiceInfo          `json:"mbsServInfo,omitempty"`
	AnyUeInd          *bool                    `json:"anyUeInd,omitempty"`
	MbsSessionId      *MbsSessionId            `json:"mbsSessionId,omitempty"`
	IngressTunAddr    []TunnelAddress          `json:"ingressTunAddr,omitempty"`
	Ssm               *Ssm                     `json:"ssm,omitempty"`
	ExtMbsServiceArea *ExternalMbsServiceArea  `json:"extMbsServiceArea,omitempty"`
	Snssai            *Snssai                  `json:"snssai,omitempty"`
	ActivationTime    string                   `json:"activationTime,omitempty"`
	MbsSessionSubsc   *MbsSessionSubscription  `json:"mbsSessionSubsc,omitempty"`
	ExpirationTime    string                   `json:"expirationTime,omitempty"`
	IngressTunAddrReq *bool                    `json:"ingressTunAddrReq,omitempty"`
	RedMbsServArea    *MbsServiceArea          `json:"redMbsServArea,omitempty"`
	Dnn               string                   `json:"dnn,omitempty"`
	TerminationTime   string                   `json:"terminationTime,omitempty"`
	ActivityStatus    MbsSessionActivityStatus `json:"activityStatus,omitempty"`
	TmgiAllocReq      *bool                    `json:"tmgiAllocReq,omitempty"`
	MbsServiceArea    *MbsServiceArea          `json:"mbsServiceArea,omitempty"`
	StartTime         string                   `json:"startTime,omitempty"`
}
