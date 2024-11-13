package models

type AppListForUeComm struct {
	OccurRatio      *int             `json:"occurRatio,omitempty"`
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty"`
	AppId           string           `json:"appId"`
	StartTime       string           `json:"startTime,omitempty"`
	AppDur          *int             `json:"appDur,omitempty"`
}
