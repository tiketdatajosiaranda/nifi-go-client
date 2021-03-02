package nifi

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"path"
)

const (
	contentTypeHeader   = "Content-Type"
	contentTypeForm     = "application/x-www-form-urlencoded"
	contentTypeJson     = "application/json"
	authorizationHeader = "Authorization"

	unknownStatus = "UNKNOWN"

	badURLError        = "bad server URL"
	badCredentialError = "username/password was not valid"
	failedMarshalError = "failed to marshal object"
)

type Context struct {
	client       *resty.Client
	baseURL      string
	userName     string
	password     string
	token        string
	Access       *Access
	Flow         *Flow
	ProcessGroup *ProcessGroup
	Processor    *Processor
	Template     *Template
}

func NewContext(s string) (*Context, error) {
	ctx := &Context{client: resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})}
	err := ctx.setBaseURL(s)
	if err != nil {
		return nil, err
	}

	ctx.Access = &Access{context: ctx}
	ctx.Flow = &Flow{context: ctx}
	ctx.ProcessGroup = &ProcessGroup{context: ctx}
	ctx.Processor = &Processor{context: ctx}
	ctx.Template = &Template{context: ctx}
	return ctx, err
}

func (c *Context) setBaseURL(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	c.baseURL = u.String()
	return nil
}

func (c *Context) GetBaseURL() string { return c.baseURL }

func (c *Context) GetURL(p string) (string, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, p)
	return u.String(), err
}

func (c *Context) getRequest(relURL string) ([]byte, error) {
	_, err := c.Access.GetStatus()
	if err != nil {
		err = c.Access.Login()
		if err != nil {
			return nil, err
		}
	}

	u, err := c.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	res, err := c.client.R().SetHeader(contentTypeHeader, contentTypeJson).Get(u)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: err}
	}

	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return nil, &HttpError{Code: res.StatusCode(), Err: fmt.Errorf(res.String())}
	}

	return res.Body(), nil
}

func (c *Context) postRequest(relURL string, body interface{}) ([]byte, error) {
	_, err := c.Access.GetStatus()
	if err != nil {
		err = c.Access.Login()
		if err != nil {
			return nil, err
		}
	}

	u, err := c.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	var res *resty.Response
	if body != nil {
		res, err = c.client.R().SetHeader(contentTypeHeader, contentTypeJson).SetBody(body).Post(u)
	} else {
		res, err = c.client.R().SetHeader(contentTypeHeader, contentTypeJson).Post(u)
	}
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: err}
	}

	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return nil, &HttpError{Code: res.StatusCode(), Err: fmt.Errorf(res.String())}
	}

	return res.Body(), nil
}

func (c *Context) putRequest(relURL string, body interface{}) ([]byte, error) {
	_, err := c.Access.GetStatus()
	if err != nil {
		err = c.Access.Login()
		if err != nil {
			return nil, err
		}
	}

	u, err := c.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	var res *resty.Response
	if body != nil {
		res, err = c.client.R().SetHeader(contentTypeHeader, contentTypeJson).SetBody(body).Put(u)
	} else {
		res, err = c.client.R().SetHeader(contentTypeHeader, contentTypeJson).Put(u)
	}
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: err}
	}

	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return nil, &HttpError{Code: res.StatusCode(), Err: fmt.Errorf(res.String())}
	}

	return res.Body(), nil
}

func (c *Context) deleteRequest(relURL string, body interface{}) ([]byte, error) {
	_, err := c.Access.GetStatus()
	if err != nil {
		err = c.Access.Login()
		if err != nil {
			return nil, err
		}
	}

	u, err := c.GetURL(relURL)
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: fmt.Errorf(badURLError)}
	}

	var res *resty.Response
	if body != nil {
		res, err = c.client.R().SetHeader(contentTypeHeader, contentTypeJson).SetBody(body).Delete(u)
	} else {
		res, err = c.client.R().SetHeader(contentTypeHeader, contentTypeJson).Delete(u)
	}
	if err != nil {
		return nil, &HttpError{Code: http.StatusInternalServerError, Err: err}
	}

	if res.StatusCode() < 200 || res.StatusCode() >= 300 {
		return nil, &HttpError{Code: res.StatusCode(), Err: fmt.Errorf(res.String())}
	}

	return res.Body(), nil
}
