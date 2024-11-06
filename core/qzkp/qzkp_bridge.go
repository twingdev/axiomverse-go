package qzkp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type QZKPResponse struct {
	Valid bool `json:"valid"`
}

func ProveKnowledge(vector []float64) (bool, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"vector": vector,
	})
	if err != nil {
		return false, err
	}

	resp, err := http.Post("http://localhost:5000/prove", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result QZKPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Valid, nil
}
