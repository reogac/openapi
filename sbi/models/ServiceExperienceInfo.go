/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Dnn              string                `json:"dnn,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	AppId            string                `json:"appId,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	Dnai             string                `json:"dnai,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
}
