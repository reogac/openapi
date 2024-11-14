package models
type ServiceExperienceInfo struct {
	 UpfInfo	*UpfInformation	`json:"upfInfo,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 NetworkArea	*NetworkAreaInfo	`json:"networkArea,omitempty"`
	 NsiId	string	`json:"nsiId,omitempty"`
	 SvcExprc	SvcExperience	`json:"svcExprc"`
	 SvcExprcVariance	*float64	`json:"svcExprcVariance,omitempty"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 SrvExpcType	ServiceExperienceType	`json:"srvExpcType,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 Ratio	*int	`json:"ratio,omitempty"`
	 RatFreq	*RatFreqInformation	`json:"ratFreq,omitempty"`
	 AppId	string	`json:"appId,omitempty"`
	 UeLocs	[]LocationInfo	`json:"ueLocs,omitempty"`
	 Dnai	string	`json:"dnai,omitempty"`
	 AppServerInst	*AddrFqdn	`json:"appServerInst,omitempty"`
	 Supis	[]string	`json:"supis,omitempty"`
}
