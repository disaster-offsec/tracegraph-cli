package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// httpClient is shared across all requests.
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// CaseCreated is returned by POST /cases.
type CaseCreated struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// CaseStatus is returned by GET /cases/{id}/status.
type CaseStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// Case is returned by GET /cases/{id}.
// Structure is not fully defined yet, so we keep it as raw JSON.
type Case map[string]any

// UploadCase sends a zip archive to POST /cases.
func UploadCase(backendURL string, zipData []byte) (*CaseCreated, error) {
	url := backendURL + "/cases"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(zipData))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/zip")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("backend returned %d: %s", resp.StatusCode, string(body))
	}

	var result CaseCreated
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}
	return &result, nil
}

// GetStatus fetches the status of a case via GET /cases/{id}/status.
func GetStatus(backendURL, caseID string) (*CaseStatus, error) {
	url := backendURL + "/cases/" + caseID + "/status"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned %d: %s", resp.StatusCode, string(body))
	}

	var result CaseStatus
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}
	return &result, nil
}

// GetCase fetches the full case via GET /cases/{id}.
func GetCase(backendURL, caseID string) (*Case, error) {
	url := backendURL + "/cases/" + caseID

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned %d: %s", resp.StatusCode, string(body))
	}

	var result Case
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}
	return &result, nil
}
