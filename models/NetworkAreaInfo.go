package models

type NetworkAreaInfo struct {
	Ecgis       []Ecgi            `json:"ecgis,omitempty"`
	Ncgis       []Ncgi            `json:"ncgis,omitempty"`
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`
	Tais        []Tai             `json:"tais,omitempty"`
}
