package models

type EventSubscription struct {
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
}
