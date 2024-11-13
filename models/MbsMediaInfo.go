package models

type MbsMediaInfo struct {
	MinReqMbsBwDl string    `json:"minReqMbsBwDl,omitempty"`
	Codecs        []string  `json:"codecs,omitempty"`
	MbsMedType    MediaType `json:"mbsMedType,omitempty"`
	MaxReqMbsBwDl string    `json:"maxReqMbsBwDl,omitempty"`
}
