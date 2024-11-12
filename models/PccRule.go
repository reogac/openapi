package models

type PccRule struct {
	ContVer         *int                               `json:"contVer,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
}
