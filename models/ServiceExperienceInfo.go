package models

type ServiceExperienceInfo struct {
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
}
