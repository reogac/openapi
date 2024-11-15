/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
}
