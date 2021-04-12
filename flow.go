package nifi

import (
	"fmt"
	"github.com/enriquebris/goconcurrentqueue"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// Flow Get the data flow, Obtain component status, Query history.
type Flow struct {
	context *Context
}

// ListTemplates Gets all templates.
// GET /flow/templates
func (f *Flow) ListTemplates() (*models.TemplatesEntity, error) {
	const relURL = "/nifi-api/flow/templates"
	raw, err := f.context.getRequest(relURL)
	if err != nil {
		return nil, err
	}

	result := models.TemplatesEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// GetProcessGroup Gets a process group.
// GET /flow/process-groups/{processGroupID}
func (f *Flow) GetProcessGroup(processGroupID string) (*models.ProcessGroupFlowEntity, error) {
	const relURL = "/nifi-api/flow/process-groups/%s"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, processGroupID))
	if err != nil {
		return nil, err
	}

	result := models.ProcessGroupFlowEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ToggleScheduleComponents Schedule or un-schedule components in the specified Process Group.
// PUT /flow/process-groups/{processGroupID}
func (f *Flow) ToggleScheduleComponents(processGroupID string, body *models.ScheduleComponentsEntity) (*models.ScheduleComponentsEntity, error) {
	const relURL = "/nifi-api/flow/process-groups/%s"
	raw, err := f.context.putRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.ScheduleComponentsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListRegistries Gets the listing of available registries.
// GET /flow/registries
func (f *Flow) ListRegistries() (*models.RegistryClientsEntity, error) {
	const relURL = "/nifi-api/flow/registries"
	raw, err := f.context.getRequest(relURL)
	if err != nil {
		return nil, err
	}

	result := models.RegistryClientsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListBuckets Gets the buckets from the specified registry for the current user.
// GET /flow/registries/{registryID}/buckets
func (f *Flow) ListBuckets(registryID string) (*models.BucketsEntity, error) {
	const relURL = "/nifi-api/flow/registries/%s/buckets"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, registryID))
	if err != nil {
		return nil, err
	}

	result := models.BucketsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListFlows Gets the flows from the specified registry and bucket for the current user.
// GET /flow/registries/{registryID}/buckets/{bucketID}/flows
func (f *Flow) ListFlows(registryID, bucketID string) (*models.VersionedFlowsEntity, error) {
	const relURL = "/nifi-api/flow/registries/%s/buckets/%s/flows"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, registryID, bucketID))
	if err != nil {
		return nil, err
	}

	result := models.VersionedFlowsEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListFlowVersions Gets the flow versions from the specified registry and bucket for the specified flow for the current user.
// GET /flow/registries/{registryID}/buckets/{bucketID}/flows/{flowID}/versions
func (f *Flow) ListFlowVersions(registryID, bucketID, flowID string) (*models.VersionedFlowSnapshotMetadataSetEntity, error) {
	const relURL = "/nifi-api/flow/registries/%s/buckets/%s/flows/%s/versions"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, registryID, bucketID, flowID))
	if err != nil {
		return nil, err
	}

	result := models.VersionedFlowSnapshotMetadataSetEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// ListProcessGroupProcessors Gets a list of process group processors.
// Function will try to iterate all processors at process group ID.
func (f *Flow) ListProcessGroupProcessors(processGroupID string) (*models.ProcessorsEntity, error) {
	pgQueue := goconcurrentqueue.NewFIFO()
	_ = pgQueue.Enqueue(processGroupID)

	processors := map[string]models.ProcessorEntity{}
	for {
		if pgQueue.GetLen() == 0 {
			break
		}

		id, err := pgQueue.Dequeue()
		if err != nil {
			continue
		}

		m, err := f.GetProcessGroup(fmt.Sprintf("%s", id))
		if err != nil {
			return nil, err
		}

		for _, p := range m.ProcessGroupFlow.Flow.Processors {
			if _, ok := processors[p.Component.Id]; !ok {
				processors[p.Component.Id] = p
			}
		}

		for _, pg := range m.ProcessGroupFlow.Flow.ProcessGroups {
			_ = pgQueue.Enqueue(pg.Id)
		}
	}

	ps := models.ProcessorsEntity{}
	for _, p := range processors {
		ps.Processors = append(ps.Processors, p)
	}
	return &ps, nil
}

// ListControllerServices Gets all controller services.
// GET /flow/process-groups/{processGroupID}/controller-services
func (f *Flow) ListControllerServices(processGroupID string) (*models.ControllerServicesEntity, error) {
	const relURL = "/nifi-api/flow/process-groups/%s/controller-services"
	raw, err := f.context.getRequest(fmt.Sprintf(relURL, processGroupID))
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
