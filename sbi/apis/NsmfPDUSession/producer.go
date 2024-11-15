/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NsmfPDUSession

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnPostSmContexts(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PostSmContextsRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandlePostSmContexts(body)

	// check for success response
	if rsp != nil {
		response.SetBody(201, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		response.SetBody(models.StatusFromPostSmContextsErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnSendMoData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	body := new(models.SendMoDataRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	ersp, prob := prod.HandleSendMoData(smContextRef, body)

	// check for defined error
	if ersp != nil {
		response.SetBody(models.StatusFromExtProblemDetails(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnPostPduSessions(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PostPduSessionsRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandlePostPduSessions(body)

	// check for success response
	if rsp != nil {
		response.SetBody(201, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		response.SetBody(models.StatusFromPostPduSessionsErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnRetrievePduSession(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	body := new(models.RetrieveData)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRetrievePduSession(pduSessionRef, body)

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

func OnTransferMoData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	body := new(models.TransferMoDataRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleTransferMoData(pduSessionRef, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnRetrieveSmContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	var body *models.SmContextRetrieveData
	if ctx.HaveRequestBody() {
		body = new(models.SmContextRetrieveData)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleRetrieveSmContext(smContextRef, body)

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

func OnUpdateSmContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	body := new(models.UpdateSmContextRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleUpdateSmContext(smContextRef, body)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		response.SetBody(models.StatusFromUpdateSmContextErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnReleaseSmContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'smContextRef'
	var smContextRef string
	smContextRef = ctx.Param("smContextRef")
	if len(smContextRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smContextRef is required"))
		return
	}

	// decode request body
	var body *models.ReleaseSmContextRequest
	if ctx.HaveRequestBody() {
		body = new(models.ReleaseSmContextRequest)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleReleaseSmContext(smContextRef, body)

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

func OnUpdatePduSession(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	body := new(models.UpdatePduSessionRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, ersp, prob := prod.HandleUpdatePduSession(pduSessionRef, body)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for defined error
	if ersp != nil {
		response.SetBody(models.StatusFromUpdatePduSessionErrorResponse(ersp), ersp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnReleasePduSession(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'pduSessionRef'
	var pduSessionRef string
	pduSessionRef = ctx.Param("pduSessionRef")
	if len(pduSessionRef) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "pduSessionRef is required"))
		return
	}

	// decode request body
	var body *models.ReleasePduSessionRequest
	if ctx.HaveRequestBody() {
		body = new(models.ReleasePduSessionRequest)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleReleasePduSession(pduSessionRef, body)

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
	HandlePostSmContexts(*models.PostSmContextsRequest) (*models.PostSmContextsResponse, *PostSmContextsErrorResponse, *models.ProblemDetails)

	HandleSendMoData(string, *models.SendMoDataRequest) (*ExtProblemDetails, *models.ProblemDetails)

	HandlePostPduSessions(*models.PostPduSessionsRequest) (*models.PostPduSessionsResponse, *PostPduSessionsErrorResponse, *models.ProblemDetails)

	HandleRetrievePduSession(string, *models.RetrieveData) (*models.RetrievedData, *models.ProblemDetails)

	HandleTransferMoData(string, *models.TransferMoDataRequest) *models.ProblemDetails

	HandleRetrieveSmContext(string, *models.SmContextRetrieveData) (*models.SmContextRetrievedData, *models.ProblemDetails)

	HandleUpdateSmContext(string, *models.UpdateSmContextRequest) (*models.UpdateSmContextResponse, *UpdateSmContextErrorResponse, *models.ProblemDetails)

	HandleReleaseSmContext(string, *models.ReleaseSmContextRequest) (*models.SmContextReleasedData, *models.ProblemDetails)

	HandleUpdatePduSession(string, *models.UpdatePduSessionRequest) (*models.UpdatePduSessionResponse, *UpdatePduSessionErrorResponse, *models.ProblemDetails)

	HandleReleasePduSession(string, *models.ReleasePduSessionRequest) (*models.ReleasePduSessionResponse, *models.ProblemDetails)
}
