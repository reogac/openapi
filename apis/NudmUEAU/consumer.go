package NudmUEAU

const (
	PATH_ROOT string = "nudm-ueau/v1"
)

// Summary: Generate authentication data for ProSe
// Description:
// Path: /:supiOrSuci/prose-security-information/generate-av
// Path Template: /%s/prose-security-information/generate-av
// Path Params: supiOrSuci
func GenerateProseAV(cli sbi.ConsumerClient, supiOrSuci string, body *models.ProSeAuthenticationInfoRequest) (rsp *models.ProSeAuthenticationInfoResult, err error) {
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
		rsp = new(ProSeAuthenticationInfoResult)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 501, 503:
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

// Summary: Generate authentication data for the UE
// Description:
// Path: /:supiOrSuci/security-information/generate-auth-data
// Path Template: /%s/security-information/generate-auth-data
// Path Params: supiOrSuci
func GenerateAuthData(cli sbi.ConsumerClient, supiOrSuci string, body *models.AuthenticationInfoRequest) (rsp *models.AuthenticationInfoResult, err error) {
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
		rsp = new(AuthenticationInfoResult)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 501, 503:
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

// Summary: Get authentication data for the FN-RG
// Description:
// Path: /:supiOrSuci/security-information-rg
// Path Template: /%s/security-information-rg
// Path Params: supiOrSuci
type GetRgAuthDataParams struct {
	SupiOrSuci        string
	AuthenticatedInd  bool
	SupportedFeatures string
	PlmnId            *PlmnId
	IfNoneMatch       string
	IfModifiedSince   string
}

func GetRgAuthData(cli sbi.ConsumerClient, params GetRgAuthDataParams) (rsp *models.RgAuthCtx, err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(RgAuthCtx)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: Create a new confirmation event
// Description:
// Path: /:supi/auth-events
// Path Template: /%s/auth-events
// Path Params: supi
func ConfirmAuth(cli sbi.ConsumerClient, supi string, body *models.AuthEvent) (rsp *models.AuthEvent, err error) {
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
		rsp = new(AuthEvent)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 404, 500, 503:
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

// Summary: Generate authentication data for the UE in EPS or IMS domain
// Description:
// Path: /:supi/hss-security-information/:hssAuthType/generate-av
// Path Template: /:supi/hss-security-information/:hssAuthType/generate-av
// Path Params: supi, hssAuthType
type GenerateAvParams struct {
	Supi        string
	HssAuthType string
}

func GenerateAv(cli sbi.ConsumerClient, params GenerateAvParams, body *models.HssAuthenticationInfoRequest) (rsp *models.HssAuthenticationInfoResult, err error) {
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
		rsp = new(HssAuthenticationInfoResult)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 501, 503:
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

// Summary: Deletes the authentication result in the UDM
// Description:
// Path: /:supi/auth-events/:authEventId
// Path Template: /:supi/auth-events/:authEventId
// Path Params: supi, authEventId
type DeleteAuthParams struct {
	Supi        string
	AuthEventId string
}

func DeleteAuth(cli sbi.ConsumerClient, params DeleteAuthParams, body *models.AuthEvent) (err error) {
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
	case 400, 404, 500, 503:
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

// Summary: Generate authentication data for the UE in GBA domain
// Description:
// Path: /:supi/gba-security-information/generate-av
// Path Template: /%s/gba-security-information/generate-av
// Path Params: supi
func GenerateGbaAv(cli sbi.ConsumerClient, supi string, body *models.GbaAuthenticationInfoRequest) (rsp *models.GbaAuthenticationInfoResult, err error) {
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
		rsp = new(GbaAuthenticationInfoResult)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 501, 503:
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
