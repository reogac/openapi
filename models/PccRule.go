package models
type PccRule struct {
	 DdNotifCtrl	*DownlinkDataNotificationControl	`json:"ddNotifCtrl,omitempty"`
	 AppReloc	*bool	`json:"appReloc,omitempty"`
	 RefChgN3gData	[]string	`json:"refChgN3gData,omitempty"`
	 ContVer	*int	`json:"contVer,omitempty"`
	 PccRuleId	string	`json:"pccRuleId"`
	 AfSigProtocol	AfSigProtocol	`json:"afSigProtocol,omitempty"`
	 RefAltQosParams	[]string	`json:"refAltQosParams,omitempty"`
	 RefTcData	[]string	`json:"refTcData,omitempty"`
	 RefChgData	[]string	`json:"refChgData,omitempty"`
	 AppId	string	`json:"appId,omitempty"`
	 AppDescriptor	string	`json:"appDescriptor,omitempty"`
	 RefQosMon	[]string	`json:"refQosMon,omitempty"`
	 DisUeNotif	*bool	`json:"disUeNotif,omitempty"`
	 PackFiltAllPrec	*int	`json:"packFiltAllPrec,omitempty"`
	 RefUmData	[]string	`json:"refUmData,omitempty"`
	 RefCondData	string	`json:"refCondData,omitempty"`
	 AddrPreserInd	*bool	`json:"addrPreserInd,omitempty"`
	 EasRedisInd	*bool	`json:"easRedisInd,omitempty"`
	 RefQosData	[]string	`json:"refQosData,omitempty"`
	 RefUmN3gData	[]string	`json:"refUmN3gData,omitempty"`
	 TscaiInputDl	*TscaiInputContainer	`json:"tscaiInputDl,omitempty"`
	 TscaiInputUl	*TscaiInputContainer	`json:"tscaiInputUl,omitempty"`
	 TscaiTimeDom	*int	`json:"tscaiTimeDom,omitempty"`
	 DdNotifCtrl2	*DownlinkDataNotificationControlRm	`json:"ddNotifCtrl2,omitempty"`
	 FlowInfos	[]FlowInformation	`json:"flowInfos,omitempty"`
	 Precedence	*int	`json:"precedence,omitempty"`
}
