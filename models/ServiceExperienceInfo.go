package models

type ServiceExperienceInfo struct {
	UeLocs           []LocationInfo      `json:"ueLocs,omitempty"`
	AppServerInst    *AddrFqdn           `json:"appServerInst,omitempty"`
	Confidence       *int                `json:"confidence,omitempty"`
	Dnn              string              `json:"dnn,omitempty"`
	Ratio            *int                `json:"ratio,omitempty"`
	RatFreq          *RatFreqInformation `json:"ratFreq,omitempty"`
	SvcExprc         SvcExperience       `json:"svcExprc"`
	Supis            []string            `json:"supis,omitempty"`
	UpfInfo          *UpfInformation     `json:"upfInfo,omitempty"`
	Dnai             string              `json:"dnai,omitempty"`
	NetworkArea      *NetworkAreaInfo    `json:"networkArea,omitempty"`
	SvcExprcVariance *float64            `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai             `json:"snssai,omitempty"`
	AppId            string              `json:"appId,omitempty"`
	SrvExpcType      string              `json:"srvExpcType,omitempty"`
	NsiId            string              `json:"nsiId,omitempty"`
}
