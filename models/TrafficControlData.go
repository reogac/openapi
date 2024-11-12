package models

type TrafficControlData struct {
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	TcId                   string                 `json:"tcId"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
}
