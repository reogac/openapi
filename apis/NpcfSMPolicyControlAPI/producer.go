package NpcfSMPolicyControlAPI
func OnCreateSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.SmPolicyContextData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleCreateSMPolicy(body)
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
func OnGetSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleGetSMPolicy()
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
func OnUpdateSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.SmPolicyUpdateContextData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleUpdateSMPolicy(body)
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
func OnDeleteSMPolicy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.SmPolicyDeleteData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
prob := HandleDeleteSMPolicy(body)
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
