package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"

	"github.com/stretchr/testify/assert"
)

type ErrorResponse struct {
    Errors []string `json:"errors"`
    Summary string `json:"summary"`
}

func TestGetValidation(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/events/i?since=tooshort&count=-1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code, "Wrong parameters fail on GET")

	responseBody := w.Body.String()

	m := ErrorResponse{}
	errUnMarshalling := json.Unmarshal([]byte(responseBody), &m)
	assert.Nil(t, errUnMarshalling)
	assert.Equal(t, 3, len(m.Errors), "Total of three failures")
	//t.Log(m)
}