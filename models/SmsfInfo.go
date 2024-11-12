package models

type SmsfInfo struct {
	SmsfInstanceId string `json:"smsfInstanceId"`
	PlmnId         PlmnId `json:"plmnId"`
	SmsfSetId      string `json:"smsfSetId,omitempty"`
}
