package proxy

import (
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

func POST(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

func PUT(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

func DELETE(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

func PATCH(w http.ResponseWriter, r *http.Request, url string, format string, token string) {
	proxyRequestWithSettings(w, r, url, format, token)
}

func proxyRequestWithSettings(w http.ResponseWriter, r* http.Request, url string, format string, token string) {
	settings := DefaultProxyRequestSettings

	settings.RequestMiddlewares = append(settings.RequestMiddlewares, PreprocessingMiddleware)
	settings.RequestMiddlewares = append(settings.RequestMiddlewares, TokenMiddlewareFactory(token))

	switch format {
	case "JSON":
		settings.ErrorHandler = JsonErrorHandler
		settings.RequestMiddlewares = append(settings.RequestMiddlewares, JsonRequestMiddlewares...)
		settings.ResponseMiddlewares = append(settings.ResponseMiddlewares, JsonResponseMiddlewares...)
	case "RAW":
		fallthrough
	default:
	}

	ProxyRequest(w, r, url, settings)
}
