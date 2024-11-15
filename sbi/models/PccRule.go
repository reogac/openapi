/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PccRule struct {
	DdNotifCtrl2    *DownlinkDataNotificationControlRm `json:"ddNotifCtrl2,omitempty"`
	DisUeNotif      *bool                              `json:"disUeNotif,omitempty"`
	PccRuleId       string                             `json:"pccRuleId"`
	RefAltQosParams []string                           `json:"refAltQosParams,omitempty"`
	RefUmN3gData    []string                           `json:"refUmN3gData,omitempty"`
	TscaiInputUl    *TscaiInputContainer               `json:"tscaiInputUl,omitempty"`
	TscaiTimeDom    *int                               `json:"tscaiTimeDom,omitempty"`
	ContVer         *int                               `json:"contVer,omitempty"`
	EasRedisInd     *bool                              `json:"easRedisInd,omitempty"`
	RefTcData       []string                           `json:"refTcData,omitempty"`
	TscaiInputDl    *TscaiInputContainer               `json:"tscaiInputDl,omitempty"`
	RefUmData       []string                           `json:"refUmData,omitempty"`
	RefQosMon       []string                           `json:"refQosMon,omitempty"`
	DdNotifCtrl     *DownlinkDataNotificationControl   `json:"ddNotifCtrl,omitempty"`
	PackFiltAllPrec *int                               `json:"packFiltAllPrec,omitempty"`
	FlowInfos       []FlowInformation                  `json:"flowInfos,omitempty"`
	Precedence      *int                               `json:"precedence,omitempty"`
	AppReloc        *bool                              `json:"appReloc,omitempty"`
	RefQosData      []string                           `json:"refQosData,omitempty"`
	RefChgN3gData   []string                           `json:"refChgN3gData,omitempty"`
	RefCondData     string                             `json:"refCondData,omitempty"`
	AddrPreserInd   *bool                              `json:"addrPreserInd,omitempty"`
	AppId           string                             `json:"appId,omitempty"`
	AppDescriptor   string                             `json:"appDescriptor,omitempty"`
	AfSigProtocol   AfSigProtocol                      `json:"afSigProtocol,omitempty"`
	RefChgData      []string                           `json:"refChgData,omitempty"`
}
