package models

type N2InfoContainer struct {
	RanInfo            *N2RanInformation `json:"ranInfo,omitempty"`
	NrppaInfo          *NrppaInformation `json:"nrppaInfo,omitempty"`
	PwsInfo            *PwsInformation   `json:"pwsInfo,omitempty"`
	V2xInfo            *V2xInformation   `json:"v2xInfo,omitempty"`
	ProseInfo          *ProSeInformation `json:"proseInfo,omitempty"`
	N2InformationClass string            `json:"n2InformationClass"`
	SmInfo             *N2SmInformation  `json:"smInfo,omitempty"`
}
