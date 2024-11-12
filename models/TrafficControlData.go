package models

type TrafficControlData struct {
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	TcId                   string                 `json:"tcId"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
}
