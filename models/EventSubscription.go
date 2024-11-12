package models

type EventSubscription struct {
	ExtraReportReq     *EventReportingRequirement    `json:"extraReportReq,omitempty"`
	MaxTopAppUlNbr     *int                          `json:"maxTopAppUlNbr,omitempty"`
	CongThresholds     []ThresholdLevel              `json:"congThresholds,omitempty"`
	RedTransReqs       []RedundantTransmissionExpReq `json:"redTransReqs,omitempty"`
	NfSetIds           []string                      `json:"nfSetIds,omitempty"`
	NfTypes            []string                      `json:"nfTypes,omitempty"`
	MaxTopAppDlNbr     *int                          `json:"maxTopAppDlNbr,omitempty"`
	NfInstanceIds      []string                      `json:"nfInstanceIds,omitempty"`
	WlanReqs           []WlanPerformanceReq          `json:"wlanReqs,omitempty"`
	UpfInfo            *UpfInformation               `json:"upfInfo,omitempty"`
	AppIds             []string                      `json:"appIds,omitempty"`
	NotificationMethod string                        `json:"notificationMethod,omitempty"`
	NfLoadLvlThds      []ThresholdLevel              `json:"nfLoadLvlThds,omitempty"`
	QosFlowRetThds     []RetainabilityThreshold      `json:"qosFlowRetThds,omitempty"`
	TgtUe              *TargetUeInformation          `json:"tgtUe,omitempty"`
	BwRequs            []BwRequirement               `json:"bwRequs,omitempty"`
	ExptAnaType        string                        `json:"exptAnaType,omitempty"`
	Dnais              []string                      `json:"dnais,omitempty"`
	LoadLevelThreshold *int                          `json:"loadLevelThreshold,omitempty"`
	VisitedAreas       []NetworkAreaInfo             `json:"visitedAreas,omitempty"`
	QosRequ            *QosRequirement               `json:"qosRequ,omitempty"`
	NwPerfRequs        []NetworkPerfRequirement      `json:"nwPerfRequs,omitempty"`
	ExptUeBehav        *ExpectedUeBehaviourData      `json:"exptUeBehav,omitempty"`
	RatFreqs           []RatFreqInformation          `json:"ratFreqs,omitempty"`
	DisperReqs         []DispersionRequirement       `json:"disperReqs,omitempty"`
	Event              string                        `json:"event"`
	LadnDnns           []string                      `json:"ladnDnns,omitempty"`
	NetworkArea        *NetworkAreaInfo              `json:"networkArea,omitempty"`
	AppServerAddrs     []AddrFqdn                    `json:"appServerAddrs,omitempty"`
	NsiLevelThrds      []int                         `json:"nsiLevelThrds,omitempty"`
	RanUeThrouThds     []string                      `json:"ranUeThrouThds,omitempty"`
	RepetitionPeriod   *int                          `json:"repetitionPeriod,omitempty"`
	Snssaia            []Snssai                      `json:"snssaia,omitempty"`
	AnySlice           *bool                         `json:"anySlice,omitempty"`
	Dnns               []string                      `json:"dnns,omitempty"`
	MatchingDir        string                        `json:"matchingDir,omitempty"`
	ExcepRequs         []Exception                   `json:"excepRequs,omitempty"`
	ListOfAnaSubsets   []string                      `json:"listOfAnaSubsets,omitempty"`
	NsiIdInfos         []NsiIdInfo                   `json:"nsiIdInfos,omitempty"`
	DnPerfReqs         []DnPerformanceReq            `json:"dnPerfReqs,omitempty"`
}
