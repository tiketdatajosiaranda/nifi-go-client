package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// FlowFileQueues View queue contents, Download flowfile content, Empty queue.
type FlowFileQueues struct {
	context *Context
}

// PostListingRequest Lists the contents of the queue in this connection.
// POST /flowfile-queues/{id}/listing-requests
func (f *FlowFileQueues) PostListingRequest(id string) (*models.ListingRequestEntity,error) {
	const relURL = "/nifi-api/flowfile-queues/%s/listing-requests"
	raw, err := f.context.postRequest(fmt.Sprintf(relURL, id), nil)
	if err != nil {
		return nil, err
	}

	result := models.ListingRequestEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// GetListingRequest Gets the current status of a listing request for the specified connection.
// GET /flowfile-queues/{id}/listing-requests/{listing-request-id}
func (f *FlowFileQueues) GetListingRequest(id,listingId string) (*models.ListingRequestEntity,error) {
	const relURL = "/nifi-api/flowfile-queues/%s/listing-requests/%s"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, id, listingId))
	if err != nil {
		return nil, err
	}

	result := models.ListingRequestEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// GetFlowFile Gets a FlowFile from a Connection.
// GET /flowfile-queues/{id}/flowFiles/{flowfile-uuid}
func (f *FlowFileQueues) GetFlowFile(id,flowfileId,clusterNodeId string) (*models.FlowFileEntity,error) {
	const relURL = "/nifi-api/flowfile-queues/%s/flowfiles/%s?clusterNodeId=%s"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, id, flowfileId,clusterNodeId))
	if err != nil {
		return nil, err
	}

	result := models.FlowFileEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}
