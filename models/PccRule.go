package models
type PccRule struct {
	 RefUmData	[]string	`json:"refUmData,omitempty"`
	 PccRuleId	string	`json:"pccRuleId"`
	 Precedence	*int	`json:"precedence,omitempty"`
	 AppReloc	*bool	`json:"appReloc,omitempty"`
	 EasRedisInd	*bool	`json:"easRedisInd,omitempty"`
	 RefAltQosParams	[]string	`json:"refAltQosParams,omitempty"`
	 RefChgData	[]string	`json:"refChgData,omitempty"`
	 RefChgN3gData	[]string	`json:"refChgN3gData,omitempty"`
	 RefQosMon	[]string	`json:"refQosMon,omitempty"`
	 TscaiInputUl	*TscaiInputContainer	`json:"tscaiInputUl,omitempty"`
	 DdNotifCtrl	*DownlinkDataNotificationControl	`json:"ddNotifCtrl,omitempty"`
	 DdNotifCtrl2	*DownlinkDataNotificationControlRm	`json:"ddNotifCtrl2,omitempty"`
	 AppDescriptor	string	`json:"appDescriptor,omitempty"`
	 ContVer	*int	`json:"contVer,omitempty"`
	 RefTcData	[]string	`json:"refTcData,omitempty"`
	 RefUmN3gData	[]string	`json:"refUmN3gData,omitempty"`
	 RefCondData	string	`json:"refCondData,omitempty"`
	 TscaiInputDl	*TscaiInputContainer	`json:"tscaiInputDl,omitempty"`
	 TscaiTimeDom	*int	`json:"tscaiTimeDom,omitempty"`
	 DisUeNotif	*bool	`json:"disUeNotif,omitempty"`
	 AfSigProtocol	AfSigProtocol	`json:"afSigProtocol,omitempty"`
	 AddrPreserInd	*bool	`json:"addrPreserInd,omitempty"`
	 FlowInfos	[]FlowInformation	`json:"flowInfos,omitempty"`
	 AppId	string	`json:"appId,omitempty"`
	 RefQosData	[]string	`json:"refQosData,omitempty"`
	 PackFiltAllPrec	*int	`json:"packFiltAllPrec,omitempty"`
}
