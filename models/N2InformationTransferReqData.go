package models

type N2InformationTransferReqData struct {
	RatSelector       RatSelector       `json:"ratSelector,omitempty"`
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`
	N2Information     N2InfoContainer   `json:"n2Information"`
	SupportedFeatures string            `json:"supportedFeatures,omitempty"`
	TaiList           []Tai             `json:"taiList,omitempty"`
}
