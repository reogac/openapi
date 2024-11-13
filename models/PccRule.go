package models

type PccRule struct {
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
}
