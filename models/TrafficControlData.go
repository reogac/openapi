package models
type TrafficControlData struct {
	 EasIpReplaceInfos	[]EasIpReplacementInfo	`json:"easIpReplaceInfos,omitempty"`
	 TcId	string	`json:"tcId"`
	 FlowStatus	FlowStatus	`json:"flowStatus,omitempty"`
	 TrafficSteeringPolIdDl	string	`json:"trafficSteeringPolIdDl,omitempty"`
	 SimConnTerm	*int	`json:"simConnTerm,omitempty"`
	 SteerModeDl	*SteeringMode	`json:"steerModeDl,omitempty"`
	 RedirectInfo	*RedirectInformation	`json:"redirectInfo,omitempty"`
	 MuteNotif	*bool	`json:"muteNotif,omitempty"`
	 SimConnInd	*bool	`json:"simConnInd,omitempty"`
	 SteerFun	SteeringFunctionality	`json:"steerFun,omitempty"`
	 AddRedirectInfo	[]RedirectInformation	`json:"addRedirectInfo,omitempty"`
	 TrafficSteeringPolIdUl	string	`json:"trafficSteeringPolIdUl,omitempty"`
	 TraffCorreInd	*bool	`json:"traffCorreInd,omitempty"`
	 SteerModeUl	*SteeringMode	`json:"steerModeUl,omitempty"`
	 MulAccCtrl	MulticastAccessControl	`json:"mulAccCtrl,omitempty"`
	 RouteToLocs	[]RouteToLocation	`json:"routeToLocs,omitempty"`
	 MaxAllowedUpLat	*int	`json:"maxAllowedUpLat,omitempty"`
	 UpPathChgEvent	*UpPathChgEvent	`json:"upPathChgEvent,omitempty"`
}
