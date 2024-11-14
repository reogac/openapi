package models
type SliceLoadLevelInformation struct {
	 LoadLevelInformation	int	`json:"loadLevelInformation"`
	 Snssais	[]Snssai	`json:"snssais"`
}
