package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

type Access struct {
	context *Context
}

func (a *Access) GetStatus() (*models.AccessStatusEntity, error) {
	const relURL = "/nifi-api/access"
	u, err := a.context.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	res, err := a.context.client.R().Get(u)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: err}
	}

	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return nil, &HttpError{Code: res.StatusCode(), Err: fmt.Errorf(res.String())}
	}

	result := models.AccessStatusEntity{}
	err = jsoniter.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	if result.AccessStatus.Status == unknownStatus {
		return nil, &HttpError{Code: http.StatusExpectationFailed, Err: fmt.Errorf(result.AccessStatus.Message)}
	}

	return &result, nil
}

func (a *Access) NewLogin(username, password string) error {
	a.context.userName = username
	a.context.password = password
	return a.Login()
}

func (a *Access) Login() error {
	if a.context.userName == "" || a.context.password == "" {
		return &HttpError{Code: http.StatusUnauthorized, Err: fmt.Errorf(badCredentialError)}
	}

	const relURL = "/nifi-api/access/token"
	u, err := a.context.GetURL(relURL)
	if err != nil {
		return &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	res, err := a.context.client.R().
		SetHeader(contentTypeHeader, contentTypeForm).
		SetBody(fmt.Sprintf("username=%s&password=%s", a.context.userName, a.context.password)).
		Post(u)
	if err != nil {
		return &HttpError{Code: http.StatusInternalServerError, Err: err}
	}

	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return &HttpError{Code: res.StatusCode(), Err: fmt.Errorf(res.String())}
	}

	token := res.String()
	a.context.client = a.context.client.SetHeader(authorizationHeader, fmt.Sprintf("Bearer %s", token))
	a.context.token = token
	return nil
}
