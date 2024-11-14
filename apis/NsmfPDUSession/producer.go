package NsmfPDUSession

func OnRetrieveSmContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.SmContextRetrieveData
	if ctx.HaveRequestBody() {
		body := new(models.SmContextRetrieveData)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleRetrieveSmContext(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnReleaseSmContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.ReleaseSmContextRequest
	if ctx.HaveRequestBody() {
		body := new(models.ReleaseSmContextRequest)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleReleaseSmContext(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnSendMoData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.SendMoDataRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		ersp, prob := HandleSendMoData(body)
		if ersp != nil {
			response.SetBody(models.StatusFromExtProblemDetails(ersp), ersp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnPostPduSessions(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.PostPduSessionsRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, ersp, prob := HandlePostPduSessions(body)
		if rsp != nil {
			response.SetBody(201, rsp)
			return
		}
		if ersp != nil {
			response.SetBody(models.StatusFromPostPduSessionsErrorResponse(ersp), ersp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnTransferMoData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.TransferMoDataRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleTransferMoData(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnPostSmContexts(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.PostSmContextsRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, ersp, prob := HandlePostSmContexts(body)
		if rsp != nil {
			response.SetBody(201, rsp)
			return
		}
		if ersp != nil {
			response.SetBody(models.StatusFromPostSmContextsErrorResponse(ersp), ersp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnUpdateSmContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.UpdateSmContextRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, ersp, prob := HandleUpdateSmContext(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if ersp != nil {
			response.SetBody(models.StatusFromUpdateSmContextErrorResponse(ersp), ersp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnUpdatePduSession(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.UpdatePduSessionRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, ersp, prob := HandleUpdatePduSession(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if ersp != nil {
			response.SetBody(models.StatusFromUpdatePduSessionErrorResponse(ersp), ersp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnReleasePduSession(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.ReleasePduSessionRequest
	if ctx.HaveRequestBody() {
		body := new(models.ReleasePduSessionRequest)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleReleasePduSession(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnRetrievePduSession(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.RetrieveData)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleRetrievePduSession(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
