package models

type TrafficControlData struct {
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	TcId                   string                 `json:"tcId"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
}
