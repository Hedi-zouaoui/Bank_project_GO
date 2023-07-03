package gapi

import (
	"net/http"
	"time"

"github.com/rs/zerolog/log"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int 
}

func (rec *ResponseRecorder) WriteHeader ( statusCode int ) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}
func HttpLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		rec := &ResponseRecorder{
			ResponseWriter: res,
			StatusCode : http.StatusOK , 
		}
		handler.ServeHTTP(res, req)
		duration := time.Since(startTime)
		logger := log.Info()

		logger.Str("protocol", "http").Str("methode", req.Method).Str("path", req.RequestURI).Int("status_code" , int(rec.StatusCode)).Str("status_text" , http.StatusText(rec.StatusCode)).Dur("duration", duration).Msg("received a HTTP request")
	})
}