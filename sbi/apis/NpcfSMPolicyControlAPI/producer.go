/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfSMPolicyControlAPI

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnCreateSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.SmPolicyContextData)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateSMPolicy(body)

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

func OnGetSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smPolicyId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleGetSMPolicy(smPolicyId)

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

func OnUpdateSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smPolicyId is required"))
		return
	}

	// decode request body
	body := new(models.SmPolicyUpdateContextData)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleUpdateSMPolicy(smPolicyId, body)

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

func OnDeleteSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'smPolicyId'
	var smPolicyId string
	smPolicyId = ctx.Param("smPolicyId")
	if len(smPolicyId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "smPolicyId is required"))
		return
	}

	// decode request body
	body := new(models.SmPolicyDeleteData)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	prob := prod.HandleDeleteSMPolicy(smPolicyId, body)

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

type Producer interface {
	HandleCreateSMPolicy(*models.SmPolicyContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleGetSMPolicy(string) (*models.SmPolicyControl, *models.ProblemDetails)

	HandleUpdateSMPolicy(string, *models.SmPolicyUpdateContextData) (*models.SmPolicyDecision, *models.ProblemDetails)

	HandleDeleteSMPolicy(string, *models.SmPolicyDeleteData) *models.ProblemDetails
}
