package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

type ProcessGroup struct {
	context *Context
}

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

func (p *ProcessGroup) DeleteProcessGroup(processGroupID string, version int64) (*models.ProcessGroupEntity, error) {
	const relURL = "/nifi-api/process-groups/%s?version=%d"
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
