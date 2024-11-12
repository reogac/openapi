package models

type PccRule struct {
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
}
