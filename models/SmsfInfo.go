package models
type SmsfInfo struct {
	 SmsfSetId	string	`json:"smsfSetId,omitempty"`
	 SmsfInstanceId	string	`json:"smsfInstanceId"`
	 PlmnId	PlmnId	`json:"plmnId"`
}
