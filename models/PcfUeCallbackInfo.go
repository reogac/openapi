package models

type PcfUeCallbackInfo struct {
	CallbackUri string   `json:"callbackUri"`
	BindingInfo []string `json:"bindingInfo,omitempty"`
}
