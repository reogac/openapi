package models

type NsiLoadLevelInfo struct {
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	NsiId                    string           `json:"nsiId,omitempty"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
}
