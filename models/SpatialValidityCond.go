package models
type SpatialValidityCond struct {
	 TrackingAreaList	[]Tai	`json:"trackingAreaList,omitempty"`
	 Countries	[]string	`json:"countries,omitempty"`
	 GeographicalServiceArea	*GeoServiceArea	`json:"geographicalServiceArea,omitempty"`
}
