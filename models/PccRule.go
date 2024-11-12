package models

type PccRule struct {
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
}
