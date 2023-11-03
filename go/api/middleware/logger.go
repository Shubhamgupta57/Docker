package middleware

import (
	"net/http"
	"time"

	"devops-testing/logger"
	"devops-testing/model"

	"github.com/felixge/httpsnoop"
)

// RequestLoggerMiddleware containing logger to log request
type RequestLoggerMiddleware struct {
	Logger *logger.Logger
}

// NewRequestLoggerMiddleware returns new request logger
func NewRequestLoggerMiddleware(logger *logger.Logger) *RequestLoggerMiddleware {
	loggerMiddleware := RequestLoggerMiddleware{
		Logger: logger,
	}
	return &loggerMiddleware
}

// GetMiddlewareHandler function returns middleware used to log requests
func (lm *RequestLoggerMiddleware) GetMiddlewareHandler() func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		metrix := httpsnoop.CaptureMetrics(next, rw, r)
		requestID := rw.Header().Get(model.HeaderRequestID)
		lm.Logger.Log.Info().
			Str("RequestID", requestID).
			Str("Host", r.Host).
			Str("Method", r.Method).
			Str("Path", r.RequestURI).
			Str("RemoteAddr", r.RemoteAddr).
			Str("Ref", r.Referer()).
			Str("UA", r.UserAgent()).
			Int("Code", metrix.Code).
			Int("Duration", int(metrix.Duration/time.Microsecond)).
			Msg("*")
	}
}
