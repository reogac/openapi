package models

type ServiceExperienceInfo struct {
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
}
