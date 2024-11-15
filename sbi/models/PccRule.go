/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	AppId           string                             `json:"appId,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
}
