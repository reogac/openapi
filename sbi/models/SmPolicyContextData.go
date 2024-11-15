/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:12 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	Pei                     string                    `json:"pei,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	Dnn                     string                    `json:"dnn"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	PduSessionId            int                       `json:"pduSessionId"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	Supi                    string                    `json:"supi"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
}
