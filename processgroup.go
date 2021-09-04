package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// ProcessGroup Create components, Instantiate a template, Upload a template.
type ProcessGroup struct {
	context *Context
}

// ListProcessGroups Gets all process groups.
// GET /process-groups/{processGroupID}/process-groups
func (p *ProcessGroup) ListProcessGroups(processGroupID string) (*models.ProcessGroupsEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/process-groups"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.ProcessGroupsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// GetProcessGroup Gets a process group.
// GET /process-groups/{processGroupID}
func (p *ProcessGroup) GetProcessGroup(processGroupID string) (*models.ProcessGroupEntity, error) {
	const relURL = "/nifi-api/process-groups/%s"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.ProcessGroupEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// CreateProcessGroup Creates a process group.
// POST /process-groups/{processGroupID}/process-groups
func (p *ProcessGroup) CreateProcessGroup(processGroupID string, body *models.ProcessGroupEntity) (*models.ProcessGroupEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/process-groups"
	raw, err := p.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.ProcessGroupEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// UpdateProcessGroup Updates a process group.
// PUT /process-groups/{processGroupID}
func (p *ProcessGroup) UpdateProcessGroup(processGroupID string, body *models.ProcessGroupEntity) (*models.ProcessGroupEntity, error) {
	const relURL = "/nifi-api/process-groups/%s"
	raw, err := p.context.putRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.ProcessGroupEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// DeleteProcessGroup Deletes a process group.
// DELETE /process-groups/{processGroupID}
func (p *ProcessGroup) DeleteProcessGroup(processGroupID string, version int64) (*models.ProcessGroupEntity, error) {
	const relURL = "/nifi-api/process-groups/%s?version=%d&disconnectedNodeAcknowledged=false"
	raw, err := p.context.deleteRequest(fmt.Sprintf(relURL, processGroupID, version), nil)
	if err != nil {
		return nil, err
	}

	result := models.ProcessGroupEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListInputPorts Gets all input ports.
// GET /process-groups/{processGroupID}/input-ports
func (p *ProcessGroup) ListInputPorts(processGroupID string) (*models.InputPortsEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/input-ports"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.InputPortsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListOutputPorts Gets all output ports.
// GET /process-groups/{processGroupID}/output-ports
func (p *ProcessGroup) ListOutputPorts(processGroupID string) (*models.OutputPortsEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/output-ports"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.OutputPortsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// CreateOutputPort Creates an output port.
// POST /process-groups/{processGroupID}/output-ports
func (p *ProcessGroup) CreateOutputPort(processGroupID string, body *models.PortEntity) (*models.PortEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/output-ports"
	raw, err := p.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.PortEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListProcessors Gets all processors.
// GET /process-groups/{processGroupID}/processors
func (p *ProcessGroup) ListProcessors(processGroupID string) (*models.ProcessorsEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/processors"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.ProcessorsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// CreateTemplateInstance Instantiates a template.
// POST /process-groups/{id}/template-instance
func (p *ProcessGroup) CreateTemplateInstance(processGroupID string, body *models.InstantiateTemplateRequestEntity) (*models.FlowEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/template-instance"
	raw, err := p.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.FlowEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListConnections Gets all connections.
// GET /process-groups/{processGroupID}/connections
func (p *ProcessGroup) ListConnections(processGroupID string) (*models.ConnectionsEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/connections"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.ConnectionsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// CreateConnection Creates a connection.
// POST /process-groups/{processGroupID}/connections
func (p *ProcessGroup) CreateConnection(processGroupID string, body *models.ConnectionEntity) (*models.ConnectionEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/connections"
	raw, err := p.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.ConnectionEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// GetVariableRegistry Gets a process group's variable registry.
// GET /process-groups/{processGroupID}/variable-registry
func (p *ProcessGroup) GetVariableRegistry(processGroupID string) (*models.VariableRegistryEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/variable-registry"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.VariableRegistryEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// GetControllerServices Gets all controller services.
// GET /flow/process-groups/{id}/controller-services
func (p *ProcessGroup) GetControllerServices(processGroupID string) (*models.ControllerServicesEntity, error) {
	const relURL = "/nifi-api/flow/process-groups/%s/controller-services"
	raw, err := p.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.ControllerServicesEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// PasteSnippet Paste copies a snippet and discards it.
// POST /process-groups/{processGroupID}/snippet-instance
func (p *ProcessGroup) PasteSnippet(processGroupID string, body *models.CopySnippetRequestEntity) (*models.FlowEntity, error) {
	const relURL = "/nifi-api/process-groups/%s/snippet-instance"
	raw, err := p.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.FlowEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}
