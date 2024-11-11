package models

type N2InformationTransferReqData struct {
	SupportedFeatures string            `json:"supportedFeatures,omitempty"`
	TaiList           []Tai             `json:"taiList,omitempty"`
	RatSelector       string            `json:"ratSelector,omitempty"`
	GlobalRanNodeList []GlobalRanNodeId `json:"globalRanNodeList,omitempty"`
	N2Information     N2InfoContainer   `json:"n2Information"`
}
