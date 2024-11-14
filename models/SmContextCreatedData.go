package models
type SmContextCreatedData struct {
	 AdditionalSnssai	*Snssai	`json:"additionalSnssai,omitempty"`
	 N2SmInfoType	N2SmInfoType	`json:"n2SmInfoType,omitempty"`
	 Gpsi	string	`json:"gpsi,omitempty"`
	 SupportedFeatures	string	`json:"supportedFeatures,omitempty"`
	 PduSessionId	*int	`json:"pduSessionId,omitempty"`
	 AllocatedEbiList	[]EbiArpMapping	`json:"allocatedEbiList,omitempty"`
	 InterPlmnApiRoot	string	`json:"interPlmnApiRoot,omitempty"`
	 SelectedSmfId	string	`json:"selectedSmfId,omitempty"`
	 SmfUri	string	`json:"smfUri,omitempty"`
	 SNssai	*Snssai	`json:"sNssai,omitempty"`
	 SmfServiceInstanceId	string	`json:"smfServiceInstanceId,omitempty"`
	 RecoveryTime	string	`json:"recoveryTime,omitempty"`
	 SelectedOldSmfId	string	`json:"selectedOldSmfId,omitempty"`
	 HSmfUri	string	`json:"hSmfUri,omitempty"`
	 UpCnxState	UpCnxState	`json:"upCnxState,omitempty"`
	 N2SmInfo	*RefToBinaryData	`json:"n2SmInfo,omitempty"`
	 HoState	HoState	`json:"hoState,omitempty"`
}
