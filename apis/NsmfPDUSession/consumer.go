package NsmfPDUSession

const (
	PATH_ROOT string = "nsmf-pdusession/v1"
)

// Summary: Retrieve SM Context
// Description:
// Path: /sm-contexts/:smContextRef/retrieve
// Path Template: /sm-contexts/%s/retrieve
// Path Params: smContextRef
func RetrieveSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.SmContextRetrieveData) (rsp *models.SmContextRetrievedData, err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmContextRetrievedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503, 504:
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

// Summary: Release SM Context
// Description:
// Path: /sm-contexts/:smContextRef/release
// Path Template: /sm-contexts/%s/release
// Path Params: smContextRef
func ReleaseSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.ReleaseSmContextRequest) (rsp *models.SmContextReleasedData, err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmContextReleasedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Send MO Data
// Description:
// Path: /sm-contexts/:smContextRef/send-mo-data
// Path Template: /sm-contexts/%s/send-mo-data
// Path Params: smContextRef
func SendMoData(cli sbi.ConsumerClient, smContextRef string, body *models.SendMoDataRequest) (ersp *models.ExtProblemDetails, err error) {
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
	case 204:
		return
	case 400, 401, 403, 404, 413, 415, 429, 500, 503:
		ersp = new(ExtProblemDetails)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411:
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

// Summary: Create
// Description:
// Path: /pdu-sessions
// Path Template: /pdu-sessions
// Path Params:
func PostPduSessions(cli sbi.ConsumerClient, body *models.PostPduSessionsRequest) (rsp *models.PostPduSessionsResponse, ersp *models.PostPduSessionsErrorResponse, err error) {
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
		rsp = new(PostPduSessionsResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
		ersp = new(PostPduSessionsErrorResponse)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411, 413, 415, 429:
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

// Summary: Transfer MO Data
// Description:
// Path: /pdu-sessions/:pduSessionRef/transfer-mo-data
// Path Template: /pdu-sessions/%s/transfer-mo-data
// Path Params: pduSessionRef
func TransferMoData(cli sbi.ConsumerClient, pduSessionRef string, body *models.TransferMoDataRequest) (err error) {
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

// Summary: Create SM Context
// Description:
// Path: /sm-contexts
// Path Template: /sm-contexts
// Path Params:
func PostSmContexts(cli sbi.ConsumerClient, body *models.PostSmContextsRequest) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse, err error) {
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
		rsp = new(PostSmContextsResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 413, 415, 429, 500, 503, 504:
		ersp = new(PostSmContextsErrorResponse)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411:
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

// Summary: Update SM Context
// Description:
// Path: /sm-contexts/:smContextRef/modify
// Path Template: /sm-contexts/%s/modify
// Path Params: smContextRef
func UpdateSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.UpdateSmContextRequest) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {
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
	case 200:
		rsp = new(UpdateSmContextResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 404, 413, 415, 429, 500, 503, 504:
		ersp = new(UpdateSmContextErrorResponse)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411:
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

// Summary: Update (initiated by V-SMF or I-SMF)
// Description:
// Path: /pdu-sessions/:pduSessionRef/modify
// Path Template: /pdu-sessions/%s/modify
// Path Params: pduSessionRef
func UpdatePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.UpdatePduSessionRequest) (rsp *models.UpdatePduSessionResponse, ersp *models.UpdatePduSessionErrorResponse, err error) {
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
	case 200:
		rsp = new(UpdatePduSessionResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 404, 500, 503:
		ersp = new(UpdatePduSessionErrorResponse)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411, 413, 415, 429:
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

// Summary: Release
// Description:
// Path: /pdu-sessions/:pduSessionRef/release
// Path Template: /pdu-sessions/%s/release
// Path Params: pduSessionRef
func ReleasePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.ReleasePduSessionRequest) (rsp *models.ReleasePduSessionResponse, err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(ReleasePduSessionResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Retrieve
// Description:
// Path: /pdu-sessions/:pduSessionRef/retrieve
// Path Template: /pdu-sessions/%s/retrieve
// Path Params: pduSessionRef
func RetrievePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.RetrieveData) (rsp *models.RetrievedData, err error) {
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
	case 200:
		rsp = new(RetrievedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503, 504:
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
