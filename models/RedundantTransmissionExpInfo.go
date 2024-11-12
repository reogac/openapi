package models

type RedundantTransmissionExpInfo struct {
	RedTransExps    []RedundantTransmissionExpPerTS `json:"redTransExps"`
	SpatialValidCon *NetworkAreaInfo                `json:"spatialValidCon,omitempty"`
	Dnn             string                          `json:"dnn,omitempty"`
}
