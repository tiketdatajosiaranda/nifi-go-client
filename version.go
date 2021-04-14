package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// Version Manage versions of process groups.
type Version struct {
	context *Context
}

// StartVersionProcessGroup Save the Process Group with the given ID.
// Begins version controlling the Process Group with the given ID or commits changes to the Versioned Flow, depending on if the provided VersionControlInformation includes a flowId.
// POST /versions/process-groups/{processGroupID}
func (v *Version) StartVersionProcessGroup(processGroupID string, body *models.StartVersionControlRequestEntity) (*models.VersionControlInformationEntity, error) {
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

// InitiateProcessGroupUpdateRequest Initiate the Update Request of a Process Group with the given ID.
// POST /versions/update-requests/process-groups/{processGroupID}
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

// GetUpdateRequest Returns the Update Request with the given ID.
// GET /versions/update-requests/{updateRequestID}
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

// DeleteUpdateRequest Deletes the Update Request with the given ID.
// DELETE /versions/update-requests/{updateRequestID}
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

// StopVersionProcessGroup Stops version controlling the Process Group with the given ID.
// DELETE /versions/process-groups/{processGroupID}
func (v *Version) StopVersionProcessGroup(processGroupID string, version int64) (*models.VersionControlInformationEntity, error) {
	const relURL = "/nifi-api/versions/process-groups/%s?disconnectedNodeAcknowledged=false&version=%d"
	raw, err := v.context.deleteRequest(fmt.Sprintf(relURL, processGroupID, version), nil)
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
