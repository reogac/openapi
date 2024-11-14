package AUSFAPI

const (
	PATH_ROOT string = "nausf-auth/v1"
)

// Summary:
// Description:
// Path: /ue-authentications/:authCtxId/5g-aka-confirmation
// Path Template: /ue-authentications/%s/5g-aka-confirmation
// Path Params: authCtxId
func Put(cli sbi.ConsumerClient, authCtxId string, body *models.ConfirmationData) (rsp *models.ConfirmationDataResponse, err error) {
	request := sbi.DefaultRequest()
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
// Path Template: /ue-authentications/%s/eap-session
// Path Params: authCtxId
func EapAuthMethod(cli sbi.ConsumerClient, authCtxId string, body *models.EapSession) (rsp *models.EapAuthMethodResponse, err error) {
	request := sbi.DefaultRequest()
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
// Path: /prose-authentications/:authCtxId/prose-auth
// Path Template: /prose-authentications/%s/prose-auth
// Path Params: authCtxId
func ProseAuth(cli sbi.ConsumerClient, authCtxId string, body *models.ProSeEapSession) (rsp *models.ProseAuthResponse, err error) {
	request := sbi.DefaultRequest()
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
// Path: /prose-authentications
// Path Template: /prose-authentications
// Path Params:
func Post(cli sbi.ConsumerClient, body *models.ProSeAuthenticationInfo) (rsp *models.ProSeAuthenticationCtx, err error) {
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
