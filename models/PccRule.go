package models

type PccRule struct {
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
}
