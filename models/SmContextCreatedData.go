package models
type SmContextCreatedData struct {
	 AdditionalSnssai	*Snssai	`json:"additionalSnssai,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 HSmfUri	string	`json:"hSmfUri,omitempty"`
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 HoState	HoState	`json:"hoState,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 SmfServiceInstanceId	string	`json:"smfServiceInstanceId,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 SmfUri	string	`json:"smfUri,omitempty"`
	 UpCnxState	UpCnxState	`json:"upCnxState,omitempty"`
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
}
