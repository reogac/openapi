package models

type EventSubscription struct {
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	MatchingDir        MatchingDirection             `json:"matchingDir,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	ExptAnaType        ExpectedAnalyticsType         `json:"exptAnaType,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	NotificationMethod NotificationMethod            `json:"notificationMethod,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	Event              NwdafEvent                    `json:"event"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
}
