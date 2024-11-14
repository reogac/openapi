/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MbsSessionSubscription struct {
	NfcInstanceId       string            `json:"nfcInstanceId,omitempty"`
	MbsSessionSubscUri  string            `json:"mbsSessionSubscUri,omitempty"`
	MbsSessionId        *MbsSessionId     `json:"mbsSessionId,omitempty"`
	AreaSessionId       *int              `json:"areaSessionId,omitempty"`
	EventList           []MbsSessionEvent `json:"eventList"`
	NotifyUri           string            `json:"notifyUri"`
	NotifyCorrelationId string            `json:"notifyCorrelationId,omitempty"`
	ExpiryTime          string            `json:"expiryTime,omitempty"`
}
