package models
type MbsMediaInfo struct {
	 MbsMedType	MediaType	`json:"mbsMedType,omitempty"`
	 MaxReqMbsBwDl	string	`json:"maxReqMbsBwDl,omitempty"`
	 MinReqMbsBwDl	string	`json:"minReqMbsBwDl,omitempty"`
	 Codecs	[]string	`json:"codecs,omitempty"`
}
