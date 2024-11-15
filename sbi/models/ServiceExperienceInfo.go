/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:28 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceExperienceInfo struct {
	Dnn              string                `json:"dnn,omitempty"`
	NetworkArea      *NetworkAreaInfo      `json:"networkArea,omitempty"`
	Ratio            *int                  `json:"ratio,omitempty"`
	SvcExprcVariance *float64              `json:"svcExprcVariance,omitempty"`
	Supis            []string              `json:"supis,omitempty"`
	Confidence       *int                  `json:"confidence,omitempty"`
	UeLocs           []LocationInfo        `json:"ueLocs,omitempty"`
	Dnai             string                `json:"dnai,omitempty"`
	SvcExprc         SvcExperience         `json:"svcExprc"`
	AppId            string                `json:"appId,omitempty"`
	NsiId            string                `json:"nsiId,omitempty"`
	AppServerInst    *AddrFqdn             `json:"appServerInst,omitempty"`
	RatFreq          *RatFreqInformation   `json:"ratFreq,omitempty"`
	Snssai           *Snssai               `json:"snssai,omitempty"`
	SrvExpcType      ServiceExperienceType `json:"srvExpcType,omitempty"`
	UpfInfo          *UpfInformation       `json:"upfInfo,omitempty"`
}
