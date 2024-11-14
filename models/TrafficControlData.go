/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficControlData struct {
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	TcId                   string                 `json:"tcId"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
}
