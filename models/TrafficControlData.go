package models

type TrafficControlData struct {
	UpPathChgEvent         *UpPathChgEvent        `json:"upPathChgEvent,omitempty"`
	TrafficSteeringPolIdDl string                 `json:"trafficSteeringPolIdDl,omitempty"`
	TrafficSteeringPolIdUl string                 `json:"trafficSteeringPolIdUl,omitempty"`
	RouteToLocs            []RouteToLocation      `json:"routeToLocs,omitempty"`
	MaxAllowedUpLat        *int                   `json:"maxAllowedUpLat,omitempty"`
	TraffCorreInd          *bool                  `json:"traffCorreInd,omitempty"`
	SimConnTerm            *int                   `json:"simConnTerm,omitempty"`
	FlowStatus             FlowStatus             `json:"flowStatus,omitempty"`
	AddRedirectInfo        []RedirectInformation  `json:"addRedirectInfo,omitempty"`
	SimConnInd             *bool                  `json:"simConnInd,omitempty"`
	SteerFun               SteeringFunctionality  `json:"steerFun,omitempty"`
	SteerModeUl            *SteeringMode          `json:"steerModeUl,omitempty"`
	MulAccCtrl             MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	TcId                   string                 `json:"tcId"`
	RedirectInfo           *RedirectInformation   `json:"redirectInfo,omitempty"`
	EasIpReplaceInfos      []EasIpReplacementInfo `json:"easIpReplaceInfos,omitempty"`
	SteerModeDl            *SteeringMode          `json:"steerModeDl,omitempty"`
	MuteNotif              *bool                  `json:"muteNotif,omitempty"`
}
