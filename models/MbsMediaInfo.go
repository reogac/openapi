package models
type MbsMediaInfo struct {
	 MbsMedType	string	`json:"mbsMedType,omitempty"`
	 MaxReqMbsBwDl	string	`json:"maxReqMbsBwDl,omitempty"`
	 MinReqMbsBwDl	string	`json:"minReqMbsBwDl,omitempty"`
	 Codecs	[]string	`json:"codecs,omitempty"`
}
