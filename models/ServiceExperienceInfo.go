package models

type ServiceExperienceInfo struct {
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
}
