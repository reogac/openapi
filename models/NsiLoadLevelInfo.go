package models

type NsiLoadLevelInfo struct {
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	NsiId                    string           `json:"nsiId,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
}
