/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SpatialValidityCond struct {
	TrackingAreaList        []Tai           `json:"trackingAreaList,omitempty"`
	Countries               []string        `json:"countries,omitempty"`
	GeographicalServiceArea *GeoServiceArea `json:"geographicalServiceArea,omitempty"`
}
