package models

type NsiLoadLevelInfo struct {
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	Snssai                   Snssai           `json:"snssai"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
	NsiId                    string           `json:"nsiId,omitempty"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
}
