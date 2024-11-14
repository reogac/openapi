package models
type ServiceExperienceInfo struct {
	 AppId	string	`json:"appId,omitempty"`
	 UeLocs	[]LocationInfo	`json:"ueLocs,omitempty"`
	 Dnai	string	`json:"dnai,omitempty"`
	 NetworkArea	*NetworkAreaInfo	`json:"networkArea,omitempty"`
	 RatFreq	*RatFreqInformation	`json:"ratFreq,omitempty"`
	 SvcExprcVariance	*float64	`json:"svcExprcVariance,omitempty"`
	 SrvExpcType	ServiceExperienceType	`json:"srvExpcType,omitempty"`
	 Confidence	*int	`json:"confidence,omitempty"`
	 Ratio	*int	`json:"ratio,omitempty"`
	 SvcExprc	SvcExperience	`json:"svcExprc"`
	 Snssai	*Snssai	`json:"snssai,omitempty"`
	 UpfInfo	*UpfInformation	`json:"upfInfo,omitempty"`
	 AppServerInst	*AddrFqdn	`json:"appServerInst,omitempty"`
	 Dnn	string	`json:"dnn,omitempty"`
	 NsiId	string	`json:"nsiId,omitempty"`
	 Supis	[]string	`json:"supis,omitempty"`
}
