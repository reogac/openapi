package models

type NetworkAreaInfo struct {
	Ncgis       []Ncgi            `json:"ncgis,omitempty"`
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`
	Tais        []Tai             `json:"tais,omitempty"`
	Ecgis       []Ecgi            `json:"ecgis,omitempty"`
}
