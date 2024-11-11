package models

type AmfStatusChangeNotification struct {
	AmfStatusInfoList []AmfStatusInfo `json:"amfStatusInfoList"`
}
