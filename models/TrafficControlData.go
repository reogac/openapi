package models
type TrafficControlData struct {
	 SteerModeUl	*SteeringMode	`json:"steerModeUl,omitempty"`
	 RedirectInfo	*RedirectInformation	`json:"redirectInfo,omitempty"`
	 MuteNotif	*bool	`json:"muteNotif,omitempty"`
	 TrafficSteeringPolIdUl	string	`json:"trafficSteeringPolIdUl,omitempty"`
	 EasIpReplaceInfos	[]EasIpReplacementInfo	`json:"easIpReplaceInfos,omitempty"`
	 SimConnTerm	*int	`json:"simConnTerm,omitempty"`
	 SteerFun	SteeringFunctionality	`json:"steerFun,omitempty"`
	 TcId	string	`json:"tcId"`
	 FlowStatus	FlowStatus	`json:"flowStatus,omitempty"`
	 AddRedirectInfo	[]RedirectInformation	`json:"addRedirectInfo,omitempty"`
	 TraffCorreInd	*bool	`json:"traffCorreInd,omitempty"`
	 SteerModeDl	*SteeringMode	`json:"steerModeDl,omitempty"`
	 TrafficSteeringPolIdDl	string	`json:"trafficSteeringPolIdDl,omitempty"`
	 RouteToLocs	[]RouteToLocation	`json:"routeToLocs,omitempty"`
	 MaxAllowedUpLat	*int	`json:"maxAllowedUpLat,omitempty"`
	 SimConnInd	*bool	`json:"simConnInd,omitempty"`
	 UpPathChgEvent	*UpPathChgEvent	`json:"upPathChgEvent,omitempty"`
	 MulAccCtrl	MulticastAccessControl	`json:"mulAccCtrl,omitempty"`
}
