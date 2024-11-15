/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package AUSFAPI

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnUeAuthentications5gAkaConfirmationPut(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// decode request body
	var body *models.ConfirmationData
	if ctx.HaveRequestBody() {
		body = new(models.ConfirmationData)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleUeAuthentications5gAkaConfirmationPut(authCtxId, body)

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

func OnEapAuthMethod(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// decode request body
	var body *models.EapSession
	if ctx.HaveRequestBody() {
		body = new(models.EapSession)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleEapAuthMethod(authCtxId, body)

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

func OnRgAuthenticationsPost(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.RgAuthenticationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleRgAuthenticationsPost(body)

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

func OnProseAuthenticationsPost(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.ProSeAuthenticationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleProseAuthenticationsPost(body)

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

func OnProseAuth(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'authCtxId'
	var authCtxId string
	authCtxId = ctx.Param("authCtxId")
	if len(authCtxId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "authCtxId is required"))
		return
	}

	// decode request body
	var body *models.ProSeEapSession
	if ctx.HaveRequestBody() {
		body = new(models.ProSeEapSession)
		if err = ctx.DecodeRequest(body); err != nil {
			response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
			return
		}
	}

	// call application handler
	rsp, prob := prod.HandleProseAuth(authCtxId, body)

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

func OnUeAuthenticationsPost(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.AuthenticationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUeAuthenticationsPost(body)

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

func OnUeAuthenticationsDeregisterPost(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.DeregistrationInfo)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleUeAuthenticationsDeregisterPost(body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

type Producer interface {
	HandleUeAuthentications5gAkaConfirmationPut(string, *models.ConfirmationData) (*models.ConfirmationDataResponse, *models.ProblemDetails)

	HandleEapAuthMethod(string, *models.EapSession) (*models.EapAuthMethodResponse, *models.ProblemDetails)

	HandleRgAuthenticationsPost(*models.RgAuthenticationInfo) (*models.RgAuthCtx, *models.ProblemDetails)

	HandleProseAuthenticationsPost(*models.ProSeAuthenticationInfo) (*models.ProSeAuthenticationCtx, *models.ProblemDetails)

	HandleProseAuth(string, *models.ProSeEapSession) (*models.ProseAuthResponse, *models.ProblemDetails)

	HandleUeAuthenticationsPost(*models.AuthenticationInfo) (*models.UEAuthenticationCtx, *models.ProblemDetails)

	HandleUeAuthenticationsDeregisterPost(*models.DeregistrationInfo) *models.ProblemDetails
}
