/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficControlData struct {
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	TcId                   string                 `json:"tcId"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
}
