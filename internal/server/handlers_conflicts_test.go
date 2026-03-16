package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashita-ai/akashi/internal/model" // Corrected module path
	"github.com/stretchr/testify/assert"
)

func TestHandleListConflicts_InvalidConflictKind(t *testing.T) {
	// Initialize handlers with a quiet logger for testing purposes
	h := &Handlers{
		logger: quietLogger(),
	}

	rec := httptest.NewRecorder()
	// Create a GET request with an invalid conflict_kind parameter to trigger validation
	req := httptest.NewRequest(http.MethodGet, "/v1/conflicts?conflict_kind=invalid_value", nil)
	h.HandleListConflicts(rec, req)

	// Verify that the server responds with 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	body := rec.Body.String()
	
	// Check if the error message correctly quotes the invalid input
	assert.Contains(t, body, `"invalid_value"`)
	
	// Verify that the response includes the list of valid conflict kinds from the model
	assert.Contains(t, body, model.ValidConflictKindsString())
	
	// Ensure the structured error code matches the expected constant
	assert.Contains(t, body, string(model.ErrCodeInvalidInput))
}
func TestHandleListConflicts_ValidConflictKind(t *testing.T) {
	h := &Handlers{
		logger: quietLogger(),
	}

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/conflicts?conflict_kind=cross_agent", nil)
	h.HandleListConflicts(rec, req)

	assert.NotEqual(t, http.StatusBadRequest, rec.Code)
}
