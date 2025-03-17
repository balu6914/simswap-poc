package database

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HarperDBClient struct {
	URL      string
	Username string
	Password string
}

func NewHarperDBClient(url, username, password string) *HarperDBClient {
	return &HarperDBClient{
		URL:      url,
		Username: username,
		Password: password,
	}
}

func (c *HarperDBClient) ExecuteOperation(operation map[string]interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(operation)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	client := &http.Client{}
	return client.Do(req)
}
