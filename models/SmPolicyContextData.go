/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:41 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	PduSessionId            int                       `json:"pduSessionId"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	Dnn                     string                    `json:"dnn"`
	NotificationUri         string                    `json:"notificationUri"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	Supi                    string                    `json:"supi"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	AccessType              AccessType                `json:"accessType,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
}
