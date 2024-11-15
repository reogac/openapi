/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:44 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EventSubscription struct {
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
}
