/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:09:23 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmUECM

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnGetNon3GppSmsfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetNon3GppSmsfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetNon3GppSmsfRegistration(&params)

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

func OnTriggerPCSCFRestoration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.TriggerRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleTriggerPCSCFRestoration(body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnSendRoutingInfoSm(ctx sbi.RequestContext, handler any) (response sbi.Response) {
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
	body := new(models.RoutingInfoSmRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleSendRoutingInfoSm(ueId, body)

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

func OnGet3GppRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params Get3GppRegistrationParams

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGet3GppRegistration(&params)

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

func OnGetSmfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

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

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetSmfRegistration(&params)

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

func OnGetNon3GppRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetNon3GppRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetNon3GppRegistration(&params)

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

func OnGetIpSmGwRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)

	// read 'ueId'
	var ueId string
	ueId = ctx.Param("ueId")
	if len(ueId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetIpSmGwRegistration(ueId)

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

func OnGetNwdafRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetNwdafRegistrationParams

	// read 'analytics-ids'
	analyticsIdsStr := ctx.Param("analytics-ids")
	if len(analyticsIdsStr) > 0 {
		var analyticsIdsTmp string
		if analyticsIdsTmp, err = models.ArrayOfStringFromString(analyticsIdsStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse analytics-ids failed: %+v", err)))
			return
		}
		params.AnalyticsIds = &analyticsIdsTmp
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetNwdafRegistration(&params)

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

func OnNwdafRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params NwdafRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'nwdafRegistrationId'
	params.NwdafRegistrationId = ctx.Param("nwdafRegistrationId")
	if len(params.NwdafRegistrationId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "nwdafRegistrationId is required"))
		return
	}

	// decode request body
	body := new(models.NwdafRegistration)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleNwdafRegistration(&params, body)

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

func OnGetRegistrations(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params GetRegistrationsParams

	// read 'dnn'
	params.Dnn = ctx.Param("dnn")

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// read 'registration-dataset-names'
	registrationDatasetNamesStr := ctx.Param("registration-dataset-names")
	if len(registrationDatasetNamesStr) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "registration-dataset-names is required"))
		return
	}

	if params.RegistrationDatasetNames, err = models.ArrayOfStringFromString(registrationDatasetNamesStr); err != nil {
		response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse registration-dataset-names failed: %+v", err)))
		return
	}

	// read 'single-nssai'
	singleNssaiStr := ctx.Param("single-nssai")
	if len(singleNssaiStr) > 0 {
		if params.SingleNssai, err = models.SnssaiFromString(singleNssaiStr); err != nil {
			response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse single-nssai failed: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleGetRegistrations(&params)

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

func OnDeregAMF(ctx sbi.RequestContext, handler any) (response sbi.Response) {
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
	body := new(models.AmfDeregInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleDeregAMF(ueId, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnPeiUpdate(ctx sbi.RequestContext, handler any) (response sbi.Response) {
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
	body := new(models.PeiUpdateInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandlePeiUpdate(ueId, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnUpdateRoamingInformation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
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
	body := new(models.RoamingInfoUpdate)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateRoamingInformation(ueId, body)

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

func OnGetLocationInfo(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params GetLocationInfoParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGetLocationInfo(&params)

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

func OnRetrieveSmfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var params RetrieveSmfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'pduSessionId'
	pduSessionIdStr := ctx.Param("pduSessionId")
	if len(pduSessionIdStr) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "pduSessionId is required"))
		return
	}

	if params.PduSessionId, err = models.IntFromString(pduSessionIdStr); err != nil {
		response.SetBody(400, models.CreateProblemDetails(400, fmt.Sprintf("parse pduSessionId failed: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRetrieveSmfRegistration(&params)

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

func OnGet3GppSmsfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var params Get3GppSmsfRegistrationParams

	// read 'ueId'
	params.UeId = ctx.Param("ueId")
	if len(params.UeId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "ueId is required"))
		return
	}

	// read 'supported-features'
	params.SupportedFeatures = ctx.Param("supported-features")

	// call application handler
	rsp, prob := prod.HandleGet3GppSmsfRegistration(&params)

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
	HandleGetNon3GppSmsfRegistration(*GetNon3GppSmsfRegistrationParams) (*models.SmsfRegistration, *models.ProblemDetails)

	HandleTriggerPCSCFRestoration(*models.TriggerRequest) *models.ProblemDetails

	HandleSendRoutingInfoSm(string, *models.RoutingInfoSmRequest) (*models.RoutingInfoSmResponse, *models.ProblemDetails)

	HandleGet3GppRegistration(*Get3GppRegistrationParams) (*models.Amf3GppAccessRegistration, *models.ProblemDetails)

	HandleGetSmfRegistration(*GetSmfRegistrationParams) (*models.SmfRegistrationInfo, *models.ProblemDetails)

	HandleGetNon3GppRegistration(*GetNon3GppRegistrationParams) (*models.AmfNon3GppAccessRegistration, *models.ProblemDetails)

	HandleGetIpSmGwRegistration(string) (*models.IpSmGwRegistration, *models.ProblemDetails)

	HandleGetNwdafRegistration(*GetNwdafRegistrationParams) (*models.NwdafRegistration, *models.ProblemDetails)

	HandleNwdafRegistration(*NwdafRegistrationParams, *models.NwdafRegistration) (*models.NwdafRegistration, *models.ProblemDetails)

	HandleGetRegistrations(*GetRegistrationsParams) (*models.RegistrationDataSets, *models.ProblemDetails)

	HandleDeregAMF(string, *models.AmfDeregInfo) *models.ProblemDetails

	HandlePeiUpdate(string, *models.PeiUpdateInfo) *models.ProblemDetails

	HandleUpdateRoamingInformation(string, *models.RoamingInfoUpdate) (*models.RoamingInfoUpdate, *models.ProblemDetails)

	HandleGetLocationInfo(*GetLocationInfoParams) (*models.LocationInfo, *models.ProblemDetails)

	HandleRetrieveSmfRegistration(*RetrieveSmfRegistrationParams) (*models.SmfRegistration, *models.ProblemDetails)

	HandleGet3GppSmsfRegistration(*Get3GppSmsfRegistrationParams) (*models.SmsfRegistration, *models.ProblemDetails)
}
