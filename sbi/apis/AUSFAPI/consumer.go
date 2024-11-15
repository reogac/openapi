/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package AUSFAPI

import (
	"sbi"
	"sbi/models"
)

const (
	PATH_ROOT string = "nausf-auth/v1"
)

// Summary:
// Description:
// Path: /prose-authentications/:authCtxId/prose-auth
// Path Params: authCtxId
func ProseAuth(cli sbi.ConsumerClient, authCtxId string, body *models.ProSeEapSession) (rsp *models.ProseAuthResponse, err error) {

	request := sbi.DefaultRequest()
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/prose-authentications/%s/prose-auth", PATH_ROOT, authCtxId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(ProseAuthResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 500:
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

// Summary:
// Description:
// Path: /ue-authentications
// Path Params:
func UeAuthenticationsPost(cli sbi.ConsumerClient, body *models.AuthenticationInfo) (rsp *models.UEAuthenticationCtx, err error) {

	request := sbi.DefaultRequest()
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-authentications", PATH_ROOT)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 201:
		rsp = new(UEAuthenticationCtx)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 501:
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

// Summary:
// Description:
// Path: /ue-authentications/deregister
// Path Params:
func UeAuthenticationsDeregisterPost(cli sbi.ConsumerClient, body *models.DeregistrationInfo) (err error) {

	request := sbi.DefaultRequest()
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-authentications/deregister", PATH_ROOT)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 204:
		return
	case 404:
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

// Summary:
// Description:
// Path: /ue-authentications/:authCtxId/5g-aka-confirmation
// Path Params: authCtxId
func UeAuthentications5gAkaConfirmationPut(cli sbi.ConsumerClient, authCtxId string, body *models.ConfirmationData) (rsp *models.ConfirmationDataResponse, err error) {

	request := sbi.DefaultRequest()
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-authentications/%s/5g-aka-confirmation", PATH_ROOT, authCtxId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(ConfirmationDataResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 500:
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

// Summary:
// Description:
// Path: /ue-authentications/:authCtxId/eap-session
// Path Params: authCtxId
func EapAuthMethod(cli sbi.ConsumerClient, authCtxId string, body *models.EapSession) (rsp *models.EapAuthMethodResponse, err error) {

	request := sbi.DefaultRequest()
	if len(authCtxId) == 0 {
		err = fmt.Errorf("authCtxId is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-authentications/%s/eap-session", PATH_ROOT, authCtxId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(EapAuthMethodResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 500:
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

// Summary:
// Description:
// Path: /rg-authentications
// Path Params:
func RgAuthenticationsPost(cli sbi.ConsumerClient, body *models.RgAuthenticationInfo) (rsp *models.RgAuthCtx, err error) {

	request := sbi.DefaultRequest()
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/rg-authentications", PATH_ROOT)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 201:
		rsp = new(RgAuthCtx)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 403, 400, 404:
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

// Summary:
// Description:
// Path: /prose-authentications
// Path Params:
func ProseAuthenticationsPost(cli sbi.ConsumerClient, body *models.ProSeAuthenticationInfo) (rsp *models.ProSeAuthenticationCtx, err error) {

	request := sbi.DefaultRequest()
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/prose-authentications", PATH_ROOT)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 201:
		rsp = new(ProSeAuthenticationCtx)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500:
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
