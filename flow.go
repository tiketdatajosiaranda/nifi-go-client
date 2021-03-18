package nifi

import (
	"fmt"
	"github.com/enriquebris/goconcurrentqueue"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

type Flow struct {
	context *Context
}

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

func (f *Flow) ListProcessGroupProcessors(processGroupID string) (*models.ProcessorsEntity, error) {
	pgQueue := goconcurrentqueue.NewFIFO()
	_ = pgQueue.Enqueue(processGroupID)

	processors := map[string]*models.ProcessorEntity{}
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
			if _, ok := processors[p.Id]; !ok {
				processors[p.Id] = &p
			}
		}

		for _, pg := range m.ProcessGroupFlow.Flow.ProcessGroups {
			_ = pgQueue.Enqueue(pg.Id)
		}
	}

	procs := models.ProcessorsEntity{}
	for _, p := range processors {
		procs.Processors = append(procs.Processors, *p)
	}
	return &procs, nil
}
