package models

type TrafficControlData struct {
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	TcId                   string                 `json:"tcId"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
}
