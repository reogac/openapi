/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfSMPolicyControlAPI

import (
	"net/http"
	"sbi"
	"sbi/models"
)

const (
	PATH_ROOT string = "npcf-smpolicycontrol/v1"
)

// Summary: Create a new Individual SM Policy
// Description:
// Path: /sm-policies
// Path Params:
func CreateSMPolicy(cli sbi.ConsumerClient, body *models.SmPolicyContextData) (rsp *models.SmPolicyDecision, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-policies", PATH_ROOT)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 201:
		rsp = new(SmPolicyDecision)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Read an Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId
// Path Params: smPolicyId
func GetSMPolicy(cli sbi.ConsumerClient, smPolicyId string) (rsp *models.SmPolicyControl, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}

	request.Path = fmt.Sprintf("%s/sm-policies/%s", PATH_ROOT, smPolicyId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(SmPolicyControl)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 429, 500, 503:
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

// Summary: Update an existing Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId/update
// Path Params: smPolicyId
func UpdateSMPolicy(cli sbi.ConsumerClient, smPolicyId string, body *models.SmPolicyUpdateContextData) (rsp *models.SmPolicyDecision, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-policies/%s/update", PATH_ROOT, smPolicyId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(SmPolicyDecision)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Delete an existing Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId/delete
// Path Params: smPolicyId
func DeleteSMPolicy(cli sbi.ConsumerClient, smPolicyId string, body *models.SmPolicyDeleteData) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(smPolicyId) == 0 {
		err = fmt.Errorf("smPolicyId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-policies/%s/delete", PATH_ROOT, smPolicyId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 204:
		return
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
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
