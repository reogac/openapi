package models

type UeContextInAmfData struct {
	EpsInterworkingInfo *EpsInterworkingInfo `json:"epsInterworkingInfo,omitempty"`
	AmfInfo             []AmfInfo            `json:"amfInfo,omitempty"`
}
