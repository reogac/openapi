/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:25 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyContextData struct {
	OnboardInd              *bool                     `json:"onboardInd,omitempty"`
	PduSessionType          PduSessionType            `json:"pduSessionType"`
	AddAccessInfo           *AdditionalAccessInfo     `json:"addAccessInfo,omitempty"`
	ServingNetwork          *PlmnIdNid                `json:"servingNetwork,omitempty"`
	Online                  *bool                     `json:"online,omitempty"`
	NwdafDatas              []NwdafData               `json:"nwdafDatas,omitempty"`
	NotificationUri         string                    `json:"notificationUri"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	Offline                 *bool                     `json:"offline,omitempty"`
	RefQosIndication        *bool                     `json:"refQosIndication,omitempty"`
	SliceInfo               Snssai                    `json:"sliceInfo"`
	QosFlowUsage            QosFlowUsage              `json:"qosFlowUsage,omitempty"`
	AtsssCapab              AtsssCapability           `json:"atsssCapab,omitempty"`
	SatBackhaulCategory     SatelliteBackhaulCategory `json:"satBackhaulCategory,omitempty"`
	AccNetChId              *AccNetChId               `json:"accNetChId,omitempty"`
	Chargingcharacteristics string                    `json:"chargingcharacteristics,omitempty"`
	IpDomain                string                    `json:"ipDomain,omitempty"`
	SubsSessAmbr            *Ambr                     `json:"subsSessAmbr,omitempty"`
	Ipv4FrameRouteList      []string                  `json:"ipv4FrameRouteList,omitempty"`
	RatType                 RatType                   `json:"ratType,omitempty"`
	ServNfId                *ServingNfIdentity        `json:"servNfId,omitempty"`
	SuppFeat                string                    `json:"suppFeat,omitempty"`
	SmfId                   string                    `json:"smfId,omitempty"`
	TraceReq                *TraceData                `json:"traceReq,omitempty"`
	PcfUeInfo               *PcfUeCallbackInfo        `json:"pcfUeInfo,omitempty"`
	InvalidSupi             *bool                     `json:"invalidSupi,omitempty"`
	DnnSelMode              DnnSelectionMode          `json:"dnnSelMode,omitempty"`
	VplmnQos                *VplmnQos                 `json:"vplmnQos,omitempty"`
	NumOfPackFilter         *int                      `json:"numOfPackFilter,omitempty"`
	Dnn                     string                    `json:"dnn"`
	RecoveryTime            string                    `json:"recoveryTime,omitempty"`
	MaPduInd                MaPduIndication           `json:"maPduInd,omitempty"`
	Ipv6FrameRouteList      []string                  `json:"ipv6FrameRouteList,omitempty"`
	AuthProfIndex           string                    `json:"authProfIndex,omitempty"`
	ThreeGppPsDataOffStatus *bool                     `json:"3gppPsDataOffStatus,omitempty"`
	PvsInfo                 []ServerAddressingInfo    `json:"pvsInfo,omitempty"`
	ChargEntityAddr         *AccNetChargingAddress    `json:"chargEntityAddr,omitempty"`
	InterGrpIds             []string                  `json:"interGrpIds,omitempty"`
	UserLocationInfo        *UserLocation             `json:"userLocationInfo,omitempty"`
	Pei                     string                    `json:"pei,omitempty"`
	Ipv4Address             string                    `json:"ipv4Address,omitempty"`
	Ipv6AddressPrefix       string                    `json:"ipv6AddressPrefix,omitempty"`
	SubsDefQos              *SubscribedDefaultQos     `json:"subsDefQos,omitempty"`
	Gpsi                    string                    `json:"gpsi,omitempty"`
	Supi                    string                    `json:"supi"`
	PduSessionId            int                       `json:"pduSessionId"`
	AccessType              AccessType                `json:"accessType,omitempty"`
}
