package AUSFAPI
func OnPost(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.ProSeAuthenticationInfo)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandlePost(body)
if rsp != nil {
response.SetBody(201, rsp)
return
}
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnPut(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
var body*models.ConfirmationData
if ctx.HaveRequestBody() {
body := new(models.ConfirmationData)
err = ctx.DecodeRequest(body)
}
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandlePut(body)
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
func OnEapAuthMethod(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
var body*models.EapSession
if ctx.HaveRequestBody() {
body := new(models.EapSession)
err = ctx.DecodeRequest(body)
}
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleEapAuthMethod(body)
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
func OnProseAuth(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
var body*models.ProSeEapSession
if ctx.HaveRequestBody() {
body := new(models.ProSeEapSession)
err = ctx.DecodeRequest(body)
}
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleProseAuth(body)
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
