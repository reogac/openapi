package NsmfPDUSession

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "RetrievePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/retrieve",
		Handler: "OnRetrievePduSession",
	},
	{
		Label:   "RetrieveSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/retrieve",
		Handler: "OnRetrieveSmContext",
	},
	{
		Label:   "ReleaseSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/release",
		Handler: "OnReleaseSmContext",
	},
	{
		Label:   "SendMoData",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/send-mo-data",
		Handler: "OnSendMoData",
	},
	{
		Label:   "PostPduSessions",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions",
		Handler: "OnPostPduSessions",
	},
	{
		Label:   "UpdatePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/modify",
		Handler: "OnUpdatePduSession",
	},
	{
		Label:   "ReleasePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/release",
		Handler: "OnReleasePduSession",
	},
	{
		Label:   "PostSmContexts",
		Method:  http.MethodPost,
		Path:    "/sm-contexts",
		Handler: "OnPostSmContexts",
	},
	{
		Label:   "UpdateSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/modify",
		Handler: "OnUpdateSmContext",
	},
	{
		Label:   "TransferMoData",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/transfer-mo-data",
		Handler: "OnTransferMoData",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "NsmfPDUSession",
		Routes:  _routes,
		Handler: p,
	}
}
