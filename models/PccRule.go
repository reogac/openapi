/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
}
