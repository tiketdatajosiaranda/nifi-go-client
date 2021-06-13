package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// Connection Create a connection, Set queue priority, Update connection destination
type Connection struct {
	context *Context
}

// GetConnection Gets a connection
// GET /connections/{id}
func (c *Connection)GetConnection(id string) (result *models.ConnectionEntity, err error) {
	const relURL = "/nifi-api/connections/%s"
	raw, err := c.context.getRequest(fmt.Sprintf(relURL, id))
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}

// PutConnection Updates a connection
// PUT /connections/{id}
func (c *Connection)PutConnection(id string,body *models.ConnectionEntity) (result *models.ConnectionEntity, err error) {
	const relURL = "/nifi-api/connections/%s"
	raw, err := c.context.putRequest(fmt.Sprintf(relURL, id),body)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}

// DeleteConnection Updates a connection
// DELETE /connections/{id}
func (c *Connection)DeleteConnection(id string,version int64) (result *models.ConnectionEntity, err error) {
	const relURL = "/nifi-api/connections/%s?version=%d&disconnectedNodeAcknowledged=false"
	raw, err := c.context.deleteRequest(fmt.Sprintf(relURL, id,version),nil)
	if err != nil {
		return
	}

	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return
}