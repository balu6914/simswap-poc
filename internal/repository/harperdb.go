package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type HarperDBClient struct {
	client *http.Client
	url    string
	user   string
	pass   string
}

func NewHarperDBClient() *HarperDBClient {
	return &HarperDBClient{
		client: &http.Client{},
		url:    os.Getenv("HARPERDB_URL"),
		user:   os.Getenv("HARPERDB_USER"),
		pass:   os.Getenv("HARPERDB_PASSWORD"),
	}
}

func (h *HarperDBClient) ExecuteQuery(payload map[string]interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", h.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(h.user, h.pass)

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(body))
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
