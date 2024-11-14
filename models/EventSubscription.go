/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:23:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
}
