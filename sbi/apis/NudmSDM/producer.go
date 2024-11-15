/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:37 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmSDM

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnGetTraceConfigData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetTraceConfigDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetTraceConfigData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetLcsMoData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetLcsMoDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetLcsMoData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetSupiOrGpsi(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSupiOrGpsiParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'mtc-provider-info'
	params.MtcProviderInfo = ctx.Param("mtc-provider-info")

	// read 'requested-gpsi-type'
	params.RequestedGpsiType = ctx.Param("requested-gpsi-type")

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'af-service-id'
	params.AfServiceId = ctx.Param("af-service-id")

	// read 'app-port-id'
	appPortIdStr := ctx.Param("app-port-id")
	if len(appPortIdStr) > 0 {
		if params.AppPortId, err = models.AppPortIdFromString(appPortIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse app-port-id failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleGetSupiOrGpsi(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetAmData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetAmDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdNidFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'adjacent-plmns'
	adjacentPlmnsStr := ctx.Param("adjacent-plmns")
	if len(adjacentPlmnsStr) > 0 {
		if params.AdjacentPlmns, err = models.ArrayOfPlmnIdFromString(adjacentPlmnsStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse adjacent-plmns failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetAmData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnUnsubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params UnsubscribeParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'subscriptionId'
	params.SubscriptionId = ctx.Param("subscriptionId")
	if len(params.SubscriptionId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribe(&params)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetIndividualSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetIndividualSharedDataParams

	// read 'sharedDataId'
	sharedDataIdStr := ctx.Param("sharedDataId")
	if len(sharedDataIdStr) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "sharedDataId is required"))
		return
	}

	if params.SharedDataId, err = models.ArrayOfStringFromString(sharedDataIdStr); err != nil {
		response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse sharedDataId failed: %+v", err)))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetIndividualSharedData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetV2xData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetV2xDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetV2xData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetMbsData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetMbsDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMbsData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetUcData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetUcDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'uc-purpose'
	params.UcPurpose = ctx.Param("uc-purpose")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// call application handler
	rsp, prob := prod.HandleGetUcData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnSorAckInfo(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleSorAckInfo(supi, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSharedDataParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'shared-data-ids'
	sharedDataIdsStr := ctx.Param("shared-data-ids")
	if len(sharedDataIdsStr) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "shared-data-ids is required"))
		return
	}

	if params.SharedDataIds, err = models.ArrayOfStringFromString(sharedDataIdsStr); err != nil {
		response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse shared-data-ids failed: %+v", err)))
		return
	}

	// read 'supportedFeatures'
	params.SupportedFeatures = ctx.Param("supportedFeatures")

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// call application handler
	rsp, prob := prod.HandleGetSharedData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetNSSAI(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetNSSAIParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetNSSAI(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetSmsMngtData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSmsMngtDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetSmsMngtData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetLcsPrivacyData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetLcsPrivacyDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetLcsPrivacyData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnUnsubscribeForSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)

	// read 'subscriptionId'
	var subscriptionId string
	subscriptionId = ctx.Param("subscriptionId")
	if len(subscriptionId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "subscriptionId is required"))
		return
	}

	// call application handler
	prob := prod.HandleUnsubscribeForSharedData(subscriptionId)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnUpuAck(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleUpuAck(supi, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetUeCtxInSmfData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetUeCtxInSmfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInSmfData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetUeCtxInSmsfData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetUeCtxInSmsfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInSmsfData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetSmData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSmDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.SnssaiFromString(singleNssaiStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)))
			return
		}
	}

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// call application handler
	rsp, prob := prod.HandleGetSmData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetSmsData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSmsDataParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSmsData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetProseData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetProseDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetProseData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetDataSets(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetDataSetsParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'dataset-names'
	datasetNamesStr := ctx.Param("dataset-names")
	if len(datasetNamesStr) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "dataset-names is required"))
		return
	}

	if params.DatasetNames, err = models.ArrayOfStringFromString(datasetNamesStr); err != nil {
		response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse dataset-names failed: %+v", err)))
		return
	}

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdNidFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetDataSets(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetUeCtxInAmfData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetUeCtxInAmfDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetUeCtxInAmfData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetEcrData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetEcrDataParams

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// call application handler
	rsp, prob := prod.HandleGetEcrData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnSNSSAIsAck(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleSNSSAIsAck(supi, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetGroupIdentifiers(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetGroupIdentifiersParams

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'ext-group-id'
	params.ExtGroupId = ctx.Param("ext-group-id")

	// read 'int-group-id'
	params.IntGroupId = ctx.Param("int-group-id")

	// read 'ue-id-ind'
	ueIdIndStr := ctx.Param("ue-id-ind")
	if len(ueIdIndStr) > 0 {
		var ueIdIndTmp bool
		if ueIdIndTmp, err = models.BoolFromString(ueIdIndStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse ue-id-ind failed: %+v", err)))
			return
		}
		params.UeIdInd = &ueIdIndTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'af-id'
	params.AfId = ctx.Param("af-id")

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// call application handler
	rsp, prob := prod.HandleGetGroupIdentifiers(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetLcsBcaData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetLcsBcaDataParams

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleGetLcsBcaData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// decode request body
	body := new(models.SdmSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSubscribe(ueId, body)

	// check for success response
	if rsp != nil {
		response.SetBody(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnSubscribeToSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.SdmSubscription)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSubscribeToSharedData(body)

	// check for success response
	if rsp != nil {
		response.SetBody(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetMultipleIdentifiers(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetMultipleIdentifiersParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'gpsi-list'
	gpsiListStr := ctx.Param("gpsi-list")
	if len(gpsiListStr) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "gpsi-list is required"))
		return
	}

	if params.GpsiList, err = models.ArrayOfStringFromString(gpsiListStr); err != nil {
		response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse gpsi-list failed: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetMultipleIdentifiers(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnGetSmfSelData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSmfSelDataParams

	// read 'plmn-id'
	plmnIdStr := ctx.Param("plmn-id")
	if len(plmnIdStr) > 0 {
		if params.PlmnId, err = models.PlmnIdFromString(plmnIdStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse plmn-id failed: %+v", err)))
			return
		}
	}

	// read 'disaster-roaming-ind'
	disasterRoamingIndStr := ctx.Param("disaster-roaming-ind")
	if len(disasterRoamingIndStr) > 0 {
		var disasterRoamingIndTmp bool
		if disasterRoamingIndTmp, err = models.BoolFromString(disasterRoamingIndStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse disaster-roaming-ind failed: %+v", err)))
			return
		}
		params.DisasterRoamingInd = &disasterRoamingIndTmp
	}

	// read 'If-None-Match'
	params.IfNoneMatch = ctx.Param("If-None-Match")

	// read 'If-Modified-Since'
	params.IfModifiedSince = ctx.Param("If-Modified-Since")

	// read 'supi'
	params.Supi = ctx.Param("supi")
	if len(params.Supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSmfSelData(&params)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnCAGAck(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body = new(models.AcknowledgeInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	prob := prod.HandleCAGAck(supi, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnUpdateSORInfo(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'supi'
	var supi string
	supi = ctx.Param("supi")
	if len(supi) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "supi is required"))
		return
	}

	// decode request body
	var body *models.SorUpdateInfo
	if ctx.HaveRequestBody() {
		body = new(models.SorUpdateInfo)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSORInfo(supi, body)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

type Producer interface {
	HandleGetTraceConfigData(*GetTraceConfigDataParams) (*models.TraceDataResponse, *models.ProblemDetails)

	HandleGetLcsMoData(*GetLcsMoDataParams) (*models.LcsMoData, *models.ProblemDetails)

	HandleGetSupiOrGpsi(*GetSupiOrGpsiParams) (*models.IdTranslationResult, *models.ProblemDetails)

	HandleGetAmData(*GetAmDataParams) (*models.AccessAndMobilitySubscriptionData, *models.ProblemDetails)

	HandleUnsubscribe(*UnsubscribeParams) *models.ProblemDetails

	HandleGetIndividualSharedData(*GetIndividualSharedDataParams) (*models.SharedData, *models.ProblemDetails)

	HandleGetV2xData(*GetV2xDataParams) (*models.V2xSubscriptionData, *models.ProblemDetails)

	HandleGetMbsData(*GetMbsDataParams) (*models.MbsSubscriptionData, *models.ProblemDetails)

	HandleGetUcData(*GetUcDataParams) (*models.UcSubscriptionData, *models.ProblemDetails)

	HandleSorAckInfo(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetSharedData(*GetSharedDataParams) (*models.SharedData, *models.ProblemDetails)

	HandleGetNSSAI(*GetNSSAIParams) (*models.Nssai, *models.ProblemDetails)

	HandleGetSmsMngtData(*GetSmsMngtDataParams) (*models.SmsManagementSubscriptionData, *models.ProblemDetails)

	HandleGetLcsPrivacyData(*GetLcsPrivacyDataParams) (*models.LcsPrivacyData, *models.ProblemDetails)

	HandleUnsubscribeForSharedData(string) *models.ProblemDetails

	HandleUpuAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetUeCtxInSmfData(*GetUeCtxInSmfDataParams) (*models.UeContextInSmfData, *models.ProblemDetails)

	HandleGetUeCtxInSmsfData(*GetUeCtxInSmsfDataParams) (*models.UeContextInSmsfData, *models.ProblemDetails)

	HandleGetSmData(*GetSmDataParams) (*models.SmSubsData, *models.ProblemDetails)

	HandleGetSmsData(*GetSmsDataParams) (*models.SmsSubscriptionData, *models.ProblemDetails)

	HandleGetProseData(*GetProseDataParams) (*models.ProseSubscriptionData, *models.ProblemDetails)

	HandleGetDataSets(*GetDataSetsParams) (*models.SubscriptionDataSets, *models.ProblemDetails)

	HandleGetUeCtxInAmfData(*GetUeCtxInAmfDataParams) (*models.UeContextInAmfData, *models.ProblemDetails)

	HandleGetEcrData(*GetEcrDataParams) (*models.EnhancedCoverageRestrictionData, *models.ProblemDetails)

	HandleSNSSAIsAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleGetGroupIdentifiers(*GetGroupIdentifiersParams) (*models.GroupIdentifiers, *models.ProblemDetails)

	HandleGetLcsBcaData(*GetLcsBcaDataParams) (*models.LcsBroadcastAssistanceTypesData, *models.ProblemDetails)

	HandleSubscribe(string, *models.SdmSubscription) (*models.SdmSubscription, *models.ProblemDetails)

	HandleSubscribeToSharedData(*models.SdmSubscription) (*models.SdmSubscription, *models.ProblemDetails)

	HandleGetMultipleIdentifiers(*GetMultipleIdentifiersParams) (*models.SupiInfo, *models.ProblemDetails)

	HandleGetSmfSelData(*GetSmfSelDataParams) (*models.SmfSelectionSubscriptionData, *models.ProblemDetails)

	HandleCAGAck(string, *models.AcknowledgeInfo) *models.ProblemDetails

	HandleUpdateSORInfo(string, *models.SorUpdateInfo) (*models.SorInfo, *models.ProblemDetails)
}
