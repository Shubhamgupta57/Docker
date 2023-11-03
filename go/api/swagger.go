package api

import (
	"devops-testing/api/wrapper"
	"devops-testing/model"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func (a *API) swaggerTest(requestCTX *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	if l := r.ContentLength; l == 0 {
		var err error
		err = errors.Wrapf(err, "Empty request data!")
		requestCTX.SetErr(err, 400)
		return
	}

	defer r.Body.Close()
	var form model.ValidateSwaggerTest

	// Converting body data into json
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		err = errors.Wrapf(err, "Unable to read json schema!")
		requestCTX.SetErr(err, 400)
		return
	}

	// Validating json data
	if errs := a.Validator.Validate(&form); errs != nil {
		requestCTX.SetErrs(errs, 400)
		return
	}

	response := model.ReturnSwaggerTest{
		Name: form.Text,
		Ball: form.Integer,
	}
	requestCTX.SetAppResponse(response, 200)
}
