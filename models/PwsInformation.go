package models
type PwsInformation struct {
	 MessageIdentifier	int	`json:"messageIdentifier"`
	 SerialNumber	int	`json:"serialNumber"`
	 PwsContainer	N2InfoContent	`json:"pwsContainer"`
	 BcEmptyAreaList	[]GlobalRanNodeId	`json:"bcEmptyAreaList,omitempty"`
	 SendRanResponse	*bool	`json:"sendRanResponse,omitempty"`
	 OmcId	string	`json:"omcId,omitempty"`
	 NfId	string	`json:"nfId,omitempty"`
}
