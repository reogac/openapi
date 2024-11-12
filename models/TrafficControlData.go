package models

type TrafficControlData struct {
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	TcId                   string                 `json:"tcId"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
}
