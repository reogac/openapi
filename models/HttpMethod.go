package models

type HttpMethod string

// Define constant values for HttpMethod
const (
	HTTPMETHOD_GET     HttpMethod = "GET"
	HTTPMETHOD_POST    HttpMethod = "POST"
	HTTPMETHOD_PUT     HttpMethod = "PUT"
	HTTPMETHOD_DELETE  HttpMethod = "DELETE"
	HTTPMETHOD_PATCH   HttpMethod = "PATCH"
	HTTPMETHOD_OPTIONS HttpMethod = "OPTIONS"
	HTTPMETHOD_HEAD    HttpMethod = "HEAD"
	HTTPMETHOD_CONNECT HttpMethod = "CONNECT"
	HTTPMETHOD_TRACE   HttpMethod = "TRACE"
)
