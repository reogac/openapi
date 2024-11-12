package models

type ServiceExperienceInfo struct {
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
}
