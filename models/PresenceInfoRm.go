package models

type PresenceInfoRm struct {
	PraId               string            `json:"praId,omitempty"`
	AdditionalPraId     string            `json:"additionalPraId,omitempty"`
	PresenceState       string            `json:"presenceState,omitempty"`
	TrackingAreaList    []Tai             `json:"trackingAreaList,omitempty"`
	EcgiList            []Ecgi            `json:"ecgiList,omitempty"`
	NcgiList            []Ncgi            `json:"ncgiList,omitempty"`
	GlobalRanNodeIdList []GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`
	GlobaleNbIdList     []GlobalRanNodeId `json:"globaleNbIdList,omitempty"`
}