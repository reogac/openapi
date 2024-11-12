package models

type EventSubscription struct {
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
}
