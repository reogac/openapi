package models

type PccRule struct {
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
}
