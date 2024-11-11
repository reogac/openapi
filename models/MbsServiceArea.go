package models

type MbsServiceArea struct {
	NcgiList []Ncgi `json:"ncgiList,omitempty"`
	TaiList  []Tai  `json:"taiList,omitempty"`
}
