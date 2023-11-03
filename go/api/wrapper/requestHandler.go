package wrapper

import (
	"devops-testing/model"
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// Request := request struct from client
type Request struct {
	HandlerFunc func(*RequestContext, http.ResponseWriter, *http.Request)
}

func (rh *Request) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	requestCTX := &RequestContext{}
	requestCTX.RequestID = uuid.NewV4().String() + "-" + uuid.NewV4().String()
	requestCTX.Path = r.URL.Path

	if requestCTX.Err == nil {
		w.Header().Set(model.HeaderRequestID, requestCTX.RequestID)
		// passing the request to respective handler
		rh.HandlerFunc(requestCTX, w, r)
	}

	switch t := requestCTX.ResponseType; t {
	case model.HTMLResp:
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(requestCTX.ResponseCode)
		w.Write(requestCTX.Response.Payload.([]byte))
	case model.JSONResp:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(requestCTX.ResponseCode)
		json.NewEncoder(w).Encode(requestCTX.Response)
	case model.ErrorResp:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(requestCTX.ResponseCode)
		requestCTX.Err.RequestID = &requestCTX.RequestID
		json.NewEncoder(w).Encode(&requestCTX.Err)
	}
}
