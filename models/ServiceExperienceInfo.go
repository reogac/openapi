package models

type ServiceExperienceInfo struct {
	Snssai           *Snssai               `json:"snssai,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Supis            []string              `json:"supis,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
}
