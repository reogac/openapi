package models

type NsiLoadLevelInfo struct {
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	NsiId                    string           `json:"nsiId,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
}
