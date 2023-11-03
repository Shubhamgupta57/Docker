package api

import (
	"devops-testing/api/wrapper"
	"net/http"
	"text/template"
)

func (a *API) landing(requestCTX *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		a.Logger.Log.Err(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	t.Execute(w, nil)
}
