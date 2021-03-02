package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

type Version struct {
	context *Context
}

func (v *Version) SaveProcessGroup(processGroupID string, body *models.StartVersionControlRequestEntity) (*models.VersionControlInformationEntity, error) {
	const relURL = "/nifi-api/versions/process-groups/%s"
	raw, err := v.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.VersionControlInformationEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

func (v *Version) InitiateProcessGroupUpdateRequest(processGroupID string, body *models.VersionControlInformationEntity) (*models.VersionedFlowUpdateRequestEntity, error) {
	const relURL = "/nifi-api/versions/update-requests/process-groups/%s"
	raw, err := v.context.postRequest(fmt.Sprintf(relURL, processGroupID), body)
	if err != nil {
		return nil, err
	}

	result := models.VersionedFlowUpdateRequestEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

func (v *Version) GetUpdateRequest(updateRequestID string) (*models.VersionedFlowUpdateRequestEntity, error) {
	const relURL = "/nifi-api/versions/update-requests/%s"
	raw, err := v.context.getRequest(fmt.Sprintf(relURL, updateRequestID))
	if err != nil {
		return nil, err
	}

	result := models.VersionedFlowUpdateRequestEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

func (v *Version) DeleteUpdateRequest(updateRequestID string) (*models.VersionedFlowUpdateRequestEntity, error) {
	const relURL = "/nifi-api/versions/update-requests/%s"
	raw, err := v.context.deleteRequest(fmt.Sprintf(relURL, updateRequestID), nil)
	if err != nil {
		return nil, err
	}

	result := models.VersionedFlowUpdateRequestEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}
