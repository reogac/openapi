/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmSDM

import (
	"net/http"
	"sbi"
	"sbi/models"
)

const (
	PATH_ROOT string = "nudm-sdm/v2"
)

// Summary: retrieve a UE's Session Management Subscription Data
// Description:
// Path: /:supi/sm-data
// Path Params: supi
type GetSmDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	SingleNssai       *Snssai
	Dnn               string
	PlmnId            *PlmnId
}

func GetSmData(cli sbi.ConsumerClient, params GetSmDataParams) (rsp *models.SmSubsData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if params.SingleNssai != nil {
		request.AddParam("single-nssai", models.SnssaiToString(*params.SingleNssai))
	}
	if len(params.Dnn) > 0 {
		request.AddParam("dnn", params.Dnn)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.Path = fmt.Sprintf("%s/%s/sm-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's LCS Privacy Subscription Data
// Description:
// Path: /:ueId/lcs-privacy-data
// Path Params: ueId
type GetLcsPrivacyDataParams struct {
	UeId              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetLcsPrivacyData(cli sbi.ConsumerClient, params GetLcsPrivacyDataParams) (rsp *models.LcsPrivacyData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/lcs-privacy-data", PATH_ROOT, params.UeId)
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

// Summary: retrieve a UE's V2X Subscription Data
// Description:
// Path: /:supi/v2x-data
// Path Params: supi
type GetV2xDataParams struct {
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
}

func GetV2xData(cli sbi.ConsumerClient, params GetV2xDataParams) (rsp *models.V2xSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/v2x-data", PATH_ROOT, params.Supi)
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

// Summary: subscribe to notifications
// Description:
// Path: /:ueId/sdm-subscriptions
// Path Params: ueId
func Subscribe(cli sbi.ConsumerClient, ueId string, body *models.SdmSubscription) (rsp *models.SdmSubscription, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/%s/sdm-subscriptions", PATH_ROOT, ueId)
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

// Summary: retrieve multiple data sets
// Description:
// Path: /:supi
// Path Params: supi
type GetDataSetsParams struct {
	IfModifiedSince    string
	Supi               string
	DatasetNames       []string
	PlmnId             *PlmnIdNid
	DisasterRoamingInd *bool
	SupportedFeatures  string
	IfNoneMatch        string
}

func GetDataSets(cli sbi.ConsumerClient, params GetDataSetsParams) (rsp *models.SubscriptionDataSets, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.DatasetNames) == 0 {
		err = fmt.Errorf("dataset-names is required")
		return
	}
	request.AddParam("dataset-names", models.ArrayOfStringToString(params.DatasetNames))
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdNidToString(*params.PlmnId))
	}
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s", PATH_ROOT, params.Supi)
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

// Summary: Mapping of UE Identifiers
// Description:
// Path: /multiple-identifiers
// Path Params:
type GetMultipleIdentifiersParams struct {
	GpsiList          []string
	SupportedFeatures string
}

func GetMultipleIdentifiers(cli sbi.ConsumerClient, params GetMultipleIdentifiersParams) (rsp *models.SupiInfo, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.GpsiList) == 0 {
		err = fmt.Errorf("gpsi-list is required")
		return
	}
	request.AddParam("gpsi-list", models.ArrayOfStringToString(params.GpsiList))
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.Path = fmt.Sprintf("%s/multiple-identifiers", PATH_ROOT)
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

// Summary: unsubscribe from notifications
// Description:
// Path: /:ueId/sdm-subscriptions/:subscriptionId
// Path Params: ueId, subscriptionId
type UnsubscribeParams struct {
	UeId           string
	SubscriptionId string
}

func Unsubscribe(cli sbi.ConsumerClient, params UnsubscribeParams) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	if len(params.SubscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	request.Path = fmt.Sprintf("%s/:ueId/sdm-subscriptions/:subscriptionId", PATH_ROOT, params.UeId, params.SubscriptionId)
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

// Summary: Nudm_Sdm Info for UPU service operation
// Description:
// Path: /:supi/am-data/upu-ack
// Path Params: supi
func UpuAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/%s/am-data/upu-ack", PATH_ROOT, supi)
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

// Summary: retrieve a UE's subscribed Enhanced Coverage Restriction Data
// Description:
// Path: /:supi/am-data/ecr-data
// Path Params: supi
type GetEcrDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetEcrData(cli sbi.ConsumerClient, params GetEcrDataParams) (rsp *models.EnhancedCoverageRestrictionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/am-data/ecr-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's UE Context In SMF Data
// Description:
// Path: /:supi/ue-context-in-smf-data
// Path Params: supi
type GetUeCtxInSmfDataParams struct {
	Supi              string
	SupportedFeatures string
}

func GetUeCtxInSmfData(cli sbi.ConsumerClient, params GetUeCtxInSmfDataParams) (rsp *models.UeContextInSmfData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.Path = fmt.Sprintf("%s/%s/ue-context-in-smf-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's SMS Subscription Data
// Description:
// Path: /:supi/sms-data
// Path Params: supi
type GetSmsDataParams struct {
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetSmsData(cli sbi.ConsumerClient, params GetSmsDataParams) (rsp *models.SmsSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.Path = fmt.Sprintf("%s/%s/sms-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's LCS Mobile Originated Subscription Data
// Description:
// Path: /:supi/lcs-mo-data
// Path Params: supi
type GetLcsMoDataParams struct {
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetLcsMoData(cli sbi.ConsumerClient, params GetLcsMoDataParams) (rsp *models.LcsMoData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/lcs-mo-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's ProSe Subscription Data
// Description:
// Path: /:supi/prose-data
// Path Params: supi
type GetProseDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
}

func GetProseData(cli sbi.ConsumerClient, params GetProseDataParams) (rsp *models.ProseSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/prose-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's SUPI or GPSI
// Description:
// Path: /:ueId/id-translation-result
// Path Params: ueId
type GetSupiOrGpsiParams struct {
	AppPortId         *AppPortId
	AfServiceId       string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
	UeId              string
	AfId              string
	MtcProviderInfo   string
	RequestedGpsiType string
}

func GetSupiOrGpsi(cli sbi.ConsumerClient, params GetSupiOrGpsiParams) (rsp *models.IdTranslationResult, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}

	if len(params.AfId) > 0 {
		request.AddParam("af-id", params.AfId)
	}
	if len(params.MtcProviderInfo) > 0 {
		request.AddParam("mtc-provider-info", params.MtcProviderInfo)
	}
	if len(params.RequestedGpsiType) > 0 {
		request.AddParam("requested-gpsi-type", params.RequestedGpsiType)
	}
	if params.AppPortId != nil {
		request.AddParam("app-port-id", models.AppPortIdToString(*params.AppPortId))
	}
	if len(params.AfServiceId) > 0 {
		request.AddParam("af-service-id", params.AfServiceId)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/id-translation-result", PATH_ROOT, params.UeId)
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

// Summary: Nudm_Sdm Info operation for S-NSSAIs acknowledgement
// Description:
// Path: /:supi/am-data/subscribed-snssais-ack
// Path Params: supi
func SNSSAIsAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/%s/am-data/subscribed-snssais-ack", PATH_ROOT, supi)
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

// Summary: retrieve the individual shared data
// Description:
// Path: /shared-data/:sharedDataId
// Path Params: sharedDataId
type GetIndividualSharedDataParams struct {
	SharedDataId      []string
	SupportedFeatures string
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetIndividualSharedData(cli sbi.ConsumerClient, params GetIndividualSharedDataParams) (rsp *models.SharedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SharedDataId) == 0 {
		err = fmt.Errorf("sharedDataId is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	request.Path = fmt.Sprintf("%s/shared-data/%s", PATH_ROOT, params.SharedDataId)
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

// Summary: retrieve a UE's UE Context In AMF Data
// Description:
// Path: /:supi/ue-context-in-amf-data
// Path Params: supi
type GetUeCtxInAmfDataParams struct {
	SupportedFeatures string
	Supi              string
}

func GetUeCtxInAmfData(cli sbi.ConsumerClient, params GetUeCtxInAmfDataParams) (rsp *models.UeContextInAmfData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	request.Path = fmt.Sprintf("%s/%s/ue-context-in-amf-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's LCS Broadcast Assistance Data Types Subscription Data
// Description:
// Path: /:supi/lcs-bca-data
// Path Params: supi
type GetLcsBcaDataParams struct {
	Supi              string
	SupportedFeatures string
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetLcsBcaData(cli sbi.ConsumerClient, params GetLcsBcaDataParams) (rsp *models.LcsBroadcastAssistanceTypesData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	request.Path = fmt.Sprintf("%s/%s/lcs-bca-data", PATH_ROOT, params.Supi)
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

// Summary: Nudm_Sdm Info operation for CAG acknowledgement
// Description:
// Path: /:supi/am-data/cag-ack
// Path Params: supi
func CAGAck(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/%s/am-data/cag-ack", PATH_ROOT, supi)
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

// Summary: Mapping of Group Identifiers
// Description:
// Path: /group-data/group-identifiers
// Path Params:
type GetGroupIdentifiersParams struct {
	IfModifiedSince   string
	ExtGroupId        string
	IntGroupId        string
	UeIdInd           *bool
	SupportedFeatures string
	AfId              string
	IfNoneMatch       string
}

func GetGroupIdentifiers(cli sbi.ConsumerClient, params GetGroupIdentifiersParams) (rsp *models.GroupIdentifiers, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if params.UeIdInd != nil {
		request.AddParam("ue-id-ind", models.BoolToString(*params.UeIdInd))
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.AfId) > 0 {
		request.AddParam("af-id", params.AfId)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.ExtGroupId) > 0 {
		request.AddParam("ext-group-id", params.ExtGroupId)
	}
	if len(params.IntGroupId) > 0 {
		request.AddParam("int-group-id", params.IntGroupId)
	}
	request.Path = fmt.Sprintf("%s/group-data/group-identifiers", PATH_ROOT)
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

// Summary: retrieve a UE's SMF Selection Subscription Data
// Description:
// Path: /:supi/smf-select-data
// Path Params: supi
type GetSmfSelDataParams struct {
	PlmnId             *PlmnId
	DisasterRoamingInd *bool
	IfNoneMatch        string
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
}

func GetSmfSelData(cli sbi.ConsumerClient, params GetSmfSelDataParams) (rsp *models.SmfSelectionSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	request.Path = fmt.Sprintf("%s/%s/smf-select-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's subscribed NSSAI
// Description:
// Path: /:supi/nssai
// Path Params: supi
type GetNSSAIParams struct {
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
	PlmnId             *PlmnId
	DisasterRoamingInd *bool
	IfNoneMatch        string
}

func GetNSSAI(cli sbi.ConsumerClient, params GetNSSAIParams) (rsp *models.Nssai, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	request.Path = fmt.Sprintf("%s/%s/nssai", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's SMS Management Subscription Data
// Description:
// Path: /:supi/sms-mng-data
// Path Params: supi
type GetSmsMngtDataParams struct {
	Supi              string
	SupportedFeatures string
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetSmsMngtData(cli sbi.ConsumerClient, params GetSmsMngtDataParams) (rsp *models.SmsManagementSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.Path = fmt.Sprintf("%s/%s/sms-mng-data", PATH_ROOT, params.Supi)
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

// Summary: Nudm_Sdm Info service operation
// Description:
// Path: /:supi/am-data/sor-ack
// Path Params: supi
func SorAckInfo(cli sbi.ConsumerClient, supi string, body *models.AcknowledgeInfo) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/%s/am-data/sor-ack", PATH_ROOT, supi)
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

// Summary: retrieve shared data
// Description:
// Path: /shared-data
// Path Params:
type GetSharedDataParams struct {
	IfNoneMatch       string
	IfModifiedSince   string
	SharedDataIds     []string
	SupportedFeatures string
	SupportedFeatures string
}

func GetSharedData(cli sbi.ConsumerClient, params GetSharedDataParams) (rsp *models.SharedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supportedFeatures", params.SupportedFeatures)
	}
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.SharedDataIds) == 0 {
		err = fmt.Errorf("shared-data-ids is required")
		return
	}
	request.AddParam("shared-data-ids", models.ArrayOfStringToString(params.SharedDataIds))
	request.Path = fmt.Sprintf("%s/shared-data", PATH_ROOT)
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

// Summary: unsubscribe from notifications for shared data
// Description:
// Path: /shared-data-subscriptions/:subscriptionId
// Path Params: subscriptionId
func UnsubscribeForSharedData(cli sbi.ConsumerClient, subscriptionId string) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	request.Path = fmt.Sprintf("%s/shared-data-subscriptions/%s", PATH_ROOT, subscriptionId)
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

// Summary: retrieve a UE's Trace Configuration Data
// Description:
// Path: /:supi/trace-data
// Path Params: supi
type GetTraceConfigDataParams struct {
	Supi              string
	SupportedFeatures string
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetTraceConfigData(cli sbi.ConsumerClient, params GetTraceConfigDataParams) (rsp *models.TraceDataResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdToString(*params.PlmnId))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	request.Path = fmt.Sprintf("%s/%s/trace-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's UE Context In SMSF Data
// Description:
// Path: /:supi/ue-context-in-smsf-data
// Path Params: supi
type GetUeCtxInSmsfDataParams struct {
	SupportedFeatures string
	Supi              string
}

func GetUeCtxInSmsfData(cli sbi.ConsumerClient, params GetUeCtxInSmsfDataParams) (rsp *models.UeContextInSmsfData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	request.Path = fmt.Sprintf("%s/%s/ue-context-in-smsf-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's 5MBS Subscription Data
// Description:
// Path: /:supi/5mbs-data
// Path Params: supi
type GetMbsDataParams struct {
	IfModifiedSince   string
	Supi              string
	SupportedFeatures string
	IfNoneMatch       string
}

func GetMbsData(cli sbi.ConsumerClient, params GetMbsDataParams) (rsp *models.MbsSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	request.Path = fmt.Sprintf("%s/%s/5mbs-data", PATH_ROOT, params.Supi)
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

// Summary: retrieve a UE's User Consent Subscription Data
// Description:
// Path: /:supi/uc-data
// Path Params: supi
type GetUcDataParams struct {
	SupportedFeatures string
	UcPurpose         string
	IfNoneMatch       string
	IfModifiedSince   string
	Supi              string
}

func GetUcData(cli sbi.ConsumerClient, params GetUcDataParams) (rsp *models.UcSubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if len(params.UcPurpose) > 0 {
		request.AddParam("uc-purpose", params.UcPurpose)
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	request.Path = fmt.Sprintf("%s/%s/uc-data", PATH_ROOT, params.Supi)
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

// Summary: Nudm_Sdm custom operation to trigger SOR info update
// Description:
// Path: /:supi/am-data/update-sor
// Path Params: supi
func UpdateSORInfo(cli sbi.ConsumerClient, supi string, body *models.SorUpdateInfo) (rsp *models.SorInfo, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/%s/am-data/update-sor", PATH_ROOT, supi)
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

// Summary: subscribe to notifications for shared data
// Description:
// Path: /shared-data-subscriptions
// Path Params:
func SubscribeToSharedData(cli sbi.ConsumerClient, body *models.SdmSubscription) (rsp *models.SdmSubscription, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/shared-data-subscriptions", PATH_ROOT)
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

// Summary: retrieve a UE's Access and Mobility Subscription Data
// Description:
// Path: /:supi/am-data
// Path Params: supi
type GetAmDataParams struct {
	IfNoneMatch        string
	IfModifiedSince    string
	Supi               string
	SupportedFeatures  string
	PlmnId             *PlmnIdNid
	AdjacentPlmns      []PlmnId
	DisasterRoamingInd *bool
}

func GetAmData(cli sbi.ConsumerClient, params GetAmDataParams) (rsp *models.AccessAndMobilitySubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.Supi) == 0 {
		err = fmt.Errorf("supi is required")
		return
	}

	if len(params.SupportedFeatures) > 0 {
		request.AddParam("supported-features", params.SupportedFeatures)
	}
	if params.PlmnId != nil {
		request.AddParam("plmn-id", models.PlmnIdNidToString(*params.PlmnId))
	}
	if len(params.AdjacentPlmns) > 0 {
		request.AddParam("adjacent-plmns", models.ArrayOfPlmnIdToString(params.AdjacentPlmns))
	}
	if params.DisasterRoamingInd != nil {
		request.AddParam("disaster-roaming-ind", models.BoolToString(*params.DisasterRoamingInd))
	}
	if len(params.IfNoneMatch) > 0 {
		request.AddHeader("If-None-Match", params.IfNoneMatch)
	}
	if len(params.IfModifiedSince) > 0 {
		request.AddHeader("If-Modified-Since", params.IfModifiedSince)
	}
	request.Path = fmt.Sprintf("%s/%s/am-data", PATH_ROOT, params.Supi)
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

/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/
