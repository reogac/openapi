package models

type PccRule struct {
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
}
