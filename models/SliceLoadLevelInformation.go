package models

type SliceLoadLevelInformation struct {
	Snssais              []Snssai `json:"snssais"`
	LoadLevelInformation int      `json:"loadLevelInformation"`
}
