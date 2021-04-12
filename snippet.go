package nifi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"net/http"
)

// Snippet Create a snippet, Move a snippet, Delete a snippet.
type Snippet struct {
	context *Context
}

// CreateSnippet Creates a snippet.
// The snippet will be automatically discarded if not used in a subsequent request after 1 minute.
// POST /snippets
func (s *Snippet) CreateSnippet(body *models.SnippetEntity) (*models.SnippetEntity, error) {
	const relURL = "/nifi-api/snippets"
	raw, err := s.context.postRequest(relURL, body)
	if err != nil {
		return nil, err
	}

	result := models.SnippetEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// MoveSnippet Move's the components in this Snippet into a new Process Group and discards the snippet.
// PUT /snippets/{snippetID}
func (s *Snippet) MoveSnippet(snippetID string, body *models.SnippetEntity) (*models.SnippetEntity, error) {
	const relURL = "/nifi-api/snippets/%s"
	raw, err := s.context.putRequest(fmt.Sprintf(relURL, snippetID), body)
	if err != nil {
		return nil, err
	}

	result := models.SnippetEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

// DiscardSnippet Deletes the components in a snippet and discards the snippet.
// DELETE /snippets/{snippetID}
func (s *Snippet) DiscardSnippet(snippetID string) (*models.SnippetEntity, error) {
	const relURL = "/nifi-api/snippets/%s?disconnectedNodeAcknowledged=false"
	raw, err := s.context.deleteRequest(fmt.Sprintf(relURL, snippetID), nil)
	if err != nil {
		return nil, err
	}

	result := models.SnippetEntity{}
	err = jsoniter.Unmarshal(raw, &result)
	if err != nil {
		return nil, &HttpError{Code: http.StatusBadRequest, Err: fmt.Errorf(failedMarshalError)}
	}

	return &result, nil
}

