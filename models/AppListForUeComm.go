package models
type AppListForUeComm struct {
	 AppDur	*int	`json:"appDur,omitempty"`
	 OccurRatio	*int	`json:"occurRatio,omitempty"`
	 SpatialValidity	*NetworkAreaInfo	`json:"spatialValidity,omitempty"`
	 AppId	string	`json:"appId"`
	 StartTime	string	`json:"startTime,omitempty"`
}
