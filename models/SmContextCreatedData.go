package models
type SmContextCreatedData struct {
	 SmfUri	string	`json:"smfUri,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 SmfServiceInstanceId	string	`json:"smfServiceInstanceId,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
	 HSmfUri	string	`json:"hSmfUri,omitempty"`
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 N2SmInfoType	string	`json:"n2SmInfoType,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 UpCnxState	string	`json:"upCnxState,omitempty"`
	 HoState	string	`json:"hoState,omitempty"`
}
