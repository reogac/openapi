package models

type PccRule struct {
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
}
