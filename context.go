package nifi

import (
	"crypto/tls"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"strings"
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

// Context Provides an access to Nifi Rest Api.
// The Rest Api provides programmatic access to command and control a NiFi instance in real time.
// Start and stop processors, monitor queues, query provenance data, and more.
// Each endpoint below includes a description, definitions of the expected input and output, potential response codes, and the authorizations required to invoke each service.
type Context struct {
	client   *resty.Client
	baseURL  string
	userName string
	password string
	token    string

	// Access User authentication and token endpoints.
	Access *Access

	// Flow Get the data flow, Obtain component status, Query history.
	Flow *Flow

	// ProcessGroup Create components, Instantiate a template, Upload a template.
	ProcessGroup *ProcessGroup

	// Processor Create a processor, Set properties, Schedule.
	Processor *Processor

	// Snippet Create a snippet, Move a snippet, Delete a snippet.
	Snippet *Snippet

	// Version Manage versions of process groups.
	Version *Version
}

// NewContext Returns a new context.
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
	ctx.Snippet = &Snippet{context: ctx}
	ctx.Version = &Version{context: ctx}
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

// GetBaseURL Returns a base URL for this Nifi context.
func (c *Context) GetBaseURL() string { return c.baseURL }

// GetURL Returns a base URL with path.
func (c *Context) GetURL(p string) (string, error) {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	b := u.String()
	if strings.HasSuffix(b, "/") {
		b = strings.TrimSuffix(b, "/")
	}

	if strings.HasPrefix(p, "/") {
		p = strings.TrimPrefix(p, "/")
	}

	return fmt.Sprintf("%s/%s", b, p), err
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
		fmt.Println(u)
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
