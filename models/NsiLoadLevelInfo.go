package models

type NsiLoadLevelInfo struct {
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	NsiId                    string           `json:"nsiId,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
}
