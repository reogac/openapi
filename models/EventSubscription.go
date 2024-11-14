/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:56:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
}
