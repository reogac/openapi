package models

type SmsfInfo struct {
	PlmnId         PlmnId `json:"plmnId"`
	SmsfSetId      string `json:"smsfSetId,omitempty"`
	SmsfInstanceId string `json:"smsfInstanceId"`
}
