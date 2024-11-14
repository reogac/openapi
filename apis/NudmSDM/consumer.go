package NudmSDM

const (
	PATH_ROOT string = "nudm-sdm/v2"
)

//Summary: Nudm_Sdm Info for UPU service operation
//Description:
//Path: /:supi/am-data/upu-ack
//Path Template: /%s/am-data/upu-ack
//Path Params: supi
func UpuAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's UE Context In SMSF Data
//Description:
//Path: /:supi/ue-context-in-smsf-data
//Path Template: /%s/ue-context-in-smsf-data
//Path Params: supi
type GetUeCtxInSmsfDataParams struct {
	SupportedFeatures string
	Supi              string
}

func GetUeCtxInSmsfData(cli sbi.ConsumerClient, params GetUeCtxInSmsfDataParams) (rsp *models.UeContextInSmsfData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(UeContextInSmsfData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: subscribe to notifications
//Description:
//Path: /:ueId/sdm-subscriptions
//Path Template: /%s/sdm-subscriptions
//Path Params: ueId
func Subscribe(cli sbi.ConsumerClient, ueId string, body *models.SdmSubscription) (rsp *models.SdmSubscription, err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 201:
		rsp = new(SdmSubscription)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 501, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's SUPI or GPSI
//Description:
//Path: /:ueId/id-translation-result
//Path Template: /%s/id-translation-result
//Path Params: ueId
type GetSupiOrGpsiParams struct {
	MtcProviderInfo   string
	SupportedFeatures string
	AfId              string
	IfNoneMatch       string
	UeId              string
	AppPortId         *AppPortId
	RequestedGpsiType string
	AfServiceId       string
	IfModifiedSince   string
}

func GetSupiOrGpsi(cli sbi.ConsumerClient, params GetSupiOrGpsiParams) (rsp *models.IdTranslationResult, err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(IdTranslationResult)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: subscribe to notifications for shared data
//Description:
//Path: /shared-data-subscriptions
//Path Template: /shared-data-subscriptions
//Path Params:
func SubscribeToSharedData(cli sbi.ConsumerClient, body *models.SdmSubscription) (rsp *models.SdmSubscription, err error) {
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 201:
		rsp = new(SdmSubscription)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's V2X Subscription Data
//Description:
//Path: /:supi/v2x-data
//Path Template: /%s/v2x-data
//Path Params: supi
type GetV2xDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetV2xData(cli sbi.ConsumerClient, params GetV2xDataParams) (rsp *models.V2xSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(V2xSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: Nudm_Sdm Info service operation
//Description:
//Path: /:supi/am-data/sor-ack
//Path Template: /%s/am-data/sor-ack
//Path Params: supi
func SorAckInfo(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: Nudm_Sdm Info operation for S-NSSAIs acknowledgement
//Description:
//Path: /:supi/am-data/subscribed-snssais-ack
//Path Template: /%s/am-data/subscribed-snssais-ack
//Path Params: supi
func SNSSAIsAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's LCS Broadcast Assistance Data Types Subscription Data
//Description:
//Path: /:supi/lcs-bca-data
//Path Template: /%s/lcs-bca-data
//Path Params: supi
type GetLcsBcaDataParams struct {
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetLcsBcaData(cli sbi.ConsumerClient, params GetLcsBcaDataParams) (rsp *models.LcsBroadcastAssistanceTypesData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(LcsBroadcastAssistanceTypesData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: Mapping of UE Identifiers
//Description:
//Path: /multiple-identifiers
//Path Template: /multiple-identifiers
//Path Params:
type GetMultipleIdentifiersParams struct {
	SupportedFeatures string
	GpsiList          []string
}

func GetMultipleIdentifiers(cli sbi.ConsumerClient, params GetMultipleIdentifiersParams) (rsp *models.SupiInfo, err error) {
	if len(gpsiList) == 0 {
		err = fmt.Errorf("gpsiList is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SupiInfo)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 429, 500, 502, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's subscribed Enhanced Coverage Restriction Data
//Description:
//Path: /:supi/am-data/ecr-data
//Path Template: /%s/am-data/ecr-data
//Path Params: supi
type GetEcrDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetEcrData(cli sbi.ConsumerClient, params GetEcrDataParams) (rsp *models.EnhancedCoverageRestrictionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(EnhancedCoverageRestrictionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's Session Management Subscription Data
//Description:
//Path: /:supi/sm-data
//Path Template: /%s/sm-data
//Path Params: supi
type GetSmDataParams struct {
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	SingleNssai       *Snssai
	Dnn               string
	PlmnId            *PlmnId
	IfNoneMatch       string
}

func GetSmData(cli sbi.ConsumerClient, params GetSmDataParams) (rsp *models.SmSubsData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmSubsData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's SMS Subscription Data
//Description:
//Path: /:supi/sms-data
//Path Template: /%s/sms-data
//Path Params: supi
type GetSmsDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	PlmnId            *PlmnId
}

func GetSmsData(cli sbi.ConsumerClient, params GetSmsDataParams) (rsp *models.SmsSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmsSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's Access and Mobility Subscription Data
//Description:
//Path: /:supi/am-data
//Path Template: /%s/am-data
//Path Params: supi
type GetAmDataParams struct {
	Supi               string
	SupportedFeatures  string
	PlmnId             *PlmnIdNid
	AdjacentPlmns      []PlmnId
	DisasterRoamingInd *bool
	IfNoneMatch        string
	IfModifiedSince    string
}

func GetAmData(cli sbi.ConsumerClient, params GetAmDataParams) (rsp *models.AccessAndMobilitySubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(AccessAndMobilitySubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's ProSe Subscription Data
//Description:
//Path: /:supi/prose-data
//Path Template: /%s/prose-data
//Path Params: supi
type GetProseDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetProseData(cli sbi.ConsumerClient, params GetProseDataParams) (rsp *models.ProseSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(ProseSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's 5MBS Subscription Data
//Description:
//Path: /:supi/5mbs-data
//Path Template: /%s/5mbs-data
//Path Params: supi
type GetMbsDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetMbsData(cli sbi.ConsumerClient, params GetMbsDataParams) (rsp *models.MbsSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(MbsSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's Trace Configuration Data
//Description:
//Path: /:supi/trace-data
//Path Template: /%s/trace-data
//Path Params: supi
type GetTraceConfigDataParams struct {
	SupportedFeatures string
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
}

func GetTraceConfigData(cli sbi.ConsumerClient, params GetTraceConfigDataParams) (rsp *models.TraceDataResponse, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(TraceDataResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's LCS Mobile Originated Subscription Data
//Description:
//Path: /:supi/lcs-mo-data
//Path Template: /%s/lcs-mo-data
//Path Params: supi
type GetLcsMoDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetLcsMoData(cli sbi.ConsumerClient, params GetLcsMoDataParams) (rsp *models.LcsMoData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(LcsMoData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: unsubscribe from notifications
//Description:
//Path: /:ueId/sdm-subscriptions/:subscriptionId
//Path Template: /:ueId/sdm-subscriptions/:subscriptionId
//Path Params: ueId, subscriptionId
type UnsubscribeParams struct {
	SubscriptionId string
	UeId           string
}

func Unsubscribe(cli sbi.ConsumerClient, params UnsubscribeParams) (err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve multiple data sets
//Description:
//Path: /:supi
//Path Template: /%s
//Path Params: supi
type GetDataSetsParams struct {
	Supi               string
	DatasetNames       []string
	PlmnId             *PlmnIdNid
	DisasterRoamingInd *bool
	SupportedFeatures  string
	IfNoneMatch        string
	IfModifiedSince    string
}

func GetDataSets(cli sbi.ConsumerClient, params GetDataSetsParams) (rsp *models.SubscriptionDataSets, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(datasetNames) == 0 {
		err = fmt.Errorf("datasetNames is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SubscriptionDataSets)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's subscribed NSSAI
//Description:
//Path: /:supi/nssai
//Path Template: /%s/nssai
//Path Params: supi
type GetNSSAIParams struct {
	Supi               string
	SupportedFeatures  string
	PlmnId             *PlmnId
	DisasterRoamingInd *bool
	IfNoneMatch        string
	IfModifiedSince    string
}

func GetNSSAI(cli sbi.ConsumerClient, params GetNSSAIParams) (rsp *models.Nssai, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(Nssai)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's UE Context In SMF Data
//Description:
//Path: /:supi/ue-context-in-smf-data
//Path Template: /%s/ue-context-in-smf-data
//Path Params: supi
type GetUeCtxInSmfDataParams struct {
	Supi              string
	SupportedFeatures string
}

func GetUeCtxInSmfData(cli sbi.ConsumerClient, params GetUeCtxInSmfDataParams) (rsp *models.UeContextInSmfData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(UeContextInSmfData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: Mapping of Group Identifiers
//Description:
//Path: /group-data/group-identifiers
//Path Template: /group-data/group-identifiers
//Path Params:
type GetGroupIdentifiersParams struct {
	SupportedFeatures string
	AfId              string
	IfNoneMatch       string
	IfModifiedSince   string
	ExtGroupId        string
	IntGroupId        string
	UeIdInd           *bool
}

func GetGroupIdentifiers(cli sbi.ConsumerClient, params GetGroupIdentifiersParams) (rsp *models.GroupIdentifiers, err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(GroupIdentifiers)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's UE Context In AMF Data
//Description:
//Path: /:supi/ue-context-in-amf-data
//Path Template: /%s/ue-context-in-amf-data
//Path Params: supi
type GetUeCtxInAmfDataParams struct {
	Supi              string
	SupportedFeatures string
}

func GetUeCtxInAmfData(cli sbi.ConsumerClient, params GetUeCtxInAmfDataParams) (rsp *models.UeContextInAmfData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(UeContextInAmfData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's SMF Selection Subscription Data
//Description:
//Path: /:supi/smf-select-data
//Path Template: /%s/smf-select-data
//Path Params: supi
type GetSmfSelDataParams struct {
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
	PlmnId             *PlmnId
	DisasterRoamingInd *bool
	IfNoneMatch        string
}

func GetSmfSelData(cli sbi.ConsumerClient, params GetSmfSelDataParams) (rsp *models.SmfSelectionSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmfSelectionSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's LCS Privacy Subscription Data
//Description:
//Path: /:ueId/lcs-privacy-data
//Path Template: /%s/lcs-privacy-data
//Path Params: ueId
type GetLcsPrivacyDataParams struct {
	IfModifiedSince   string
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
}

func GetLcsPrivacyData(cli sbi.ConsumerClient, params GetLcsPrivacyDataParams) (rsp *models.LcsPrivacyData, err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(LcsPrivacyData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: Nudm_Sdm custom operation to trigger SOR info update
//Description:
//Path: /:supi/am-data/update-sor
//Path Template: /%s/am-data/update-sor
//Path Params: supi
func UpdateSORInfo(cli sbi.ConsumerClient, supi string, body *models.SorUpdateInfo) (rsp *models.SorInfo, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SorInfo)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: unsubscribe from notifications for shared data
//Description:
//Path: /shared-data-subscriptions/:subscriptionId
//Path Template: /shared-data-subscriptions/%s
//Path Params: subscriptionId
func UnsubscribeForSharedData(cli sbi.ConsumerClient, subscriptionId string) (err error) {
	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's SMS Management Subscription Data
//Description:
//Path: /:supi/sms-mng-data
//Path Template: /%s/sms-mng-data
//Path Params: supi
type GetSmsMngtDataParams struct {
	Supi              string
	SupportedFeatures string
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSmsMngtData(cli sbi.ConsumerClient, params GetSmsMngtDataParams) (rsp *models.SmsManagementSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmsManagementSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve a UE's User Consent Subscription Data
//Description:
//Path: /:supi/uc-data
//Path Template: /%s/uc-data
//Path Params: supi
type GetUcDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	UcPurpose         string
}

func GetUcData(cli sbi.ConsumerClient, params GetUcDataParams) (rsp *models.UcSubscriptionData, err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(UcSubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: Nudm_Sdm Info operation for CAG acknowledgement
//Description:
//Path: /:supi/am-data/cag-ack
//Path Template: /%s/am-data/cag-ack
//Path Params: supi
func CAGAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve shared data
//Description:
//Path: /shared-data
//Path Template: /shared-data
//Path Params:
type GetSharedDataParams struct {
	SharedDataIds     []string
	SupportedFeatures string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSharedData(cli sbi.ConsumerClient, params GetSharedDataParams) (rsp *models.SharedData, err error) {
	if len(sharedDataIds) == 0 {
		err = fmt.Errorf("sharedDataIds is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SharedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

//Summary: retrieve the individual shared data
//Description:
//Path: /shared-data/:sharedDataId
//Path Template: /shared-data/%s
//Path Params: sharedDataId
type GetIndividualSharedDataParams struct {
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	SharedDataId      []string
}

func GetIndividualSharedData(cli sbi.ConsumerClient, params GetIndividualSharedDataParams) (rsp *models.SharedData, err error) {
	if len(sharedDataId) == 0 {
		err = fmt.Errorf("sharedDataId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SharedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}
