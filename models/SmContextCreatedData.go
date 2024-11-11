type SmContextCreatedData struct {
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 SmfUri	string	`json:"smfUri,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 HSmfUri	string	`json:"hSmfUri,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 UpCnxState	string	`json:"upCnxState,omitempty"`
	 N2SmInfoType	string	`json:"n2SmInfoType,omitempty"`
	 HoState	string	`json:"hoState,omitempty"`
	 SmfServiceInstanceId	string	`json:"smfServiceInstanceId,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
}
