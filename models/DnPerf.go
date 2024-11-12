package models
type DnPerf struct {
	 SpatialValidCon	*NetworkAreaInfo	`json:"spatialValidCon,omitempty"`
	 TemporalValidCon	*TimeWindow	`json:"temporalValidCon,omitempty"`
	 AppServerInsAddr	*AddrFqdn	`json:"appServerInsAddr,omitempty"`
	 UpfInfo	*UpfInformation	`json:"upfInfo,omitempty"`
	 Dnai	string	`json:"dnai,omitempty"`
	 PerfData	PerfData	`json:"perfData"`
}
