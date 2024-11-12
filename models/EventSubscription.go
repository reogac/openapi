package models

type EventSubscription struct {
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
}
