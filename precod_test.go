package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerStatus(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(
		t,
		http.StatusOK,
		responseRecorder.Code,
		fmt.Sprintf("expected status code: %d, got %d", http.StatusOK, responseRecorder.Code),
	)

	assert.NotEmpty(
		t,
		responseRecorder.Body,
		"expected body: Should NOT be empty, got empty",
	)
}

func TestMainHandlerSupportedCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=sarapul", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(
		t,
		http.StatusBadRequest,
		responseRecorder.Code,
		fmt.Sprintf("expected status code: %d, got %d", http.StatusBadRequest, responseRecorder.Code),
	)

	assert.Equal(
		t,
		"wrong city value",
		responseRecorder.Body.String(),
		fmt.Sprintf("expected body: wrong city value, got %s", responseRecorder.Body.String()),
	)

}

func TestMainHandlerCountCafe(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(
		t,
		http.StatusOK,
		responseRecorder.Code,
		fmt.Sprintf("expected status code: %d, got %d", http.StatusOK, responseRecorder.Code),
	)

	cityes := strings.Split(responseRecorder.Body.String(), ",")

	assert.Len(
		t,
		cityes,
		totalCount,
	)
}
