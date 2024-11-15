/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NsmfPDUSession

import (
	"net/http"
	"sbi"
	"sbi/models"
)

const (
	PATH_ROOT string = "nsmf-pdusession/v1"
)

// Summary: Create SM Context
// Description:
// Path: /sm-contexts
// Path Params:
func PostSmContexts(cli sbi.ConsumerClient, body *models.PostSmContextsRequest) (rsp *models.PostSmContextsResponse, ersp *models.PostSmContextsErrorResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-contexts", PATH_ROOT)
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

// Summary: Send MO Data
// Description:
// Path: /sm-contexts/:smContextRef/send-mo-data
// Path Params: smContextRef
func SendMoData(cli sbi.ConsumerClient, smContextRef string, body *models.SendMoDataRequest) (ersp *models.ExtProblemDetails, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-contexts/%s/send-mo-data", PATH_ROOT, smContextRef)
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
// Path Params:
func PostPduSessions(cli sbi.ConsumerClient, body *models.PostPduSessionsRequest) (rsp *models.PostPduSessionsResponse, ersp *models.PostPduSessionsErrorResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/pdu-sessions", PATH_ROOT)
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

// Summary: Retrieve
// Description:
// Path: /pdu-sessions/:pduSessionRef/retrieve
// Path Params: pduSessionRef
func RetrievePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.RetrieveData) (rsp *models.RetrievedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/pdu-sessions/%s/retrieve", PATH_ROOT, pduSessionRef)
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

// Summary: Transfer MO Data
// Description:
// Path: /pdu-sessions/:pduSessionRef/transfer-mo-data
// Path Params: pduSessionRef
func TransferMoData(cli sbi.ConsumerClient, pduSessionRef string, body *models.TransferMoDataRequest) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/pdu-sessions/%s/transfer-mo-data", PATH_ROOT, pduSessionRef)
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

// Summary: Retrieve SM Context
// Description:
// Path: /sm-contexts/:smContextRef/retrieve
// Path Params: smContextRef
func RetrieveSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.SmContextRetrieveData) (rsp *models.SmContextRetrievedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-contexts/%s/retrieve", PATH_ROOT, smContextRef)
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

// Summary: Update SM Context
// Description:
// Path: /sm-contexts/:smContextRef/modify
// Path Params: smContextRef
func UpdateSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.UpdateSmContextRequest) (rsp *models.UpdateSmContextResponse, ersp *models.UpdateSmContextErrorResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-contexts/%s/modify", PATH_ROOT, smContextRef)
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

// Summary: Release SM Context
// Description:
// Path: /sm-contexts/:smContextRef/release
// Path Params: smContextRef
func ReleaseSmContext(cli sbi.ConsumerClient, smContextRef string, body *models.ReleaseSmContextRequest) (rsp *models.SmContextReleasedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(smContextRef) == 0 {
		err = fmt.Errorf("smContextRef is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/sm-contexts/%s/release", PATH_ROOT, smContextRef)
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

// Summary: Update (initiated by V-SMF or I-SMF)
// Description:
// Path: /pdu-sessions/:pduSessionRef/modify
// Path Params: pduSessionRef
func UpdatePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.UpdatePduSessionRequest) (rsp *models.UpdatePduSessionResponse, ersp *models.UpdatePduSessionErrorResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/pdu-sessions/%s/modify", PATH_ROOT, pduSessionRef)
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
// Path Params: pduSessionRef
func ReleasePduSession(cli sbi.ConsumerClient, pduSessionRef string, body *models.ReleasePduSessionRequest) (rsp *models.ReleasePduSessionResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(pduSessionRef) == 0 {
		err = fmt.Errorf("pduSessionRef is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/pdu-sessions/%s/release", PATH_ROOT, pduSessionRef)
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
