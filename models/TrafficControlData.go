package models

type TrafficControlData struct {
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	TcId                   string                 `json:"tcId"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
}
