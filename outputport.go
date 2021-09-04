package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// OutputPort Create an output port, Set remote port access control
type OutputPort struct {
	context *Context
}

// GetOutputPort Gets an output port
// GET /output-ports/{id}
func (o *OutputPort)GetOutputPort(id string) (result *models.PortEntity, err error) {
	const relURL = "/nifi-api/output-ports/%s"
	raw, err := o.context.getRequest(fmt.Sprintf(relURL, id))
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}

// PutOutputPort Updates an output port
// PUT /output-ports/{id}
func (o *OutputPort)PutOutputPort(id string,body *models.PortEntity) (result *models.PortEntity, err error) {
	const relURL = "/nifi-api/output-ports/%s"
	raw, err := o.context.putRequest(fmt.Sprintf(relURL, id),body)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}

// DeleteConnection Deletes an output port
// DELETE /output-ports/{id}
func (o *OutputPort)DeleteOutputPort(id string,version int64) (result *models.PortEntity, err error) {
	const relURL = "/nifi-api/output-ports/%s?version=%d&disconnectedNodeAcknowledged=false"
	raw, err := o.context.deleteRequest(fmt.Sprintf(relURL, id,version),nil)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}

// PutOutputPortRunStatus Updates run status of an output-port
// PUT /output-ports/{id}
func (o *OutputPort)PutOutputPortRunStatus(id string,body *models.PortRunStatusEntity) (result *models.PortEntity, err error) {
	const relURL = "/nifi-api/output-ports/%s/run-status"
	raw, err := o.context.putRequest(fmt.Sprintf(relURL, id),body)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}