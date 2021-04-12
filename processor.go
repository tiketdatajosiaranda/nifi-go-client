package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// Processor Create a processor, Set properties, Schedule.
type Processor struct {
	context *Context
}

// GetProcessor Gets a processor.
// GET /processors/{processorID}
func (p *Processor) GetProcessor(processorID string) (*models.ProcessorEntity, error) {
	const relURL = "/nifi-api/processors/%s"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processorID))
	if err != nil {
		return nil, err
	}

	result := models.ProcessorEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// UpdateProcessor Updates a processor.
// PUT /processors/{processorID}
func (p *Processor) UpdateProcessor(processorID string, body *models.ProcessorEntity) (*models.ProcessorEntity, error) {
	const relURL = "/nifi-api/processors/%s"
	raw, err := p.context.putRequest(fmt.Sprintf(relURL, processorID), body)
	if err != nil {
		return nil, err
	}

	result := models.ProcessorEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// DeleteProcessor Deletes a processor.
// DELETE /processors/{processorID}
func (p *Processor) DeleteProcessor(processorID string, version int64) (*models.ProcessorEntity, error) {
	const relURL = "/nifi-api/processors/%s?version=%d"
	raw, err := p.context.deleteRequest(fmt.Sprintf(relURL, processorID, version), nil)
	if err != nil {
		return nil, err
	}

	result := models.ProcessorEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}
