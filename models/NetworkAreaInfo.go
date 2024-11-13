package models

type NetworkAreaInfo struct {
	Tais        []Tai             `json:"tais,omitempty"`
	Ecgis       []Ecgi            `json:"ecgis,omitempty"`
	Ncgis       []Ncgi            `json:"ncgis,omitempty"`
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`
}
