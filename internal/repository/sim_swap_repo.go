package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"simswap-poc/internal/domain"
)

// Define the ErrRecordNotFound error
var ErrRecordNotFound = errors.New("record not found")

type SimSwapRepository interface {
	GetSimSwapDate(phoneNumber string) (*domain.SimSwapInfo, error)
	CheckSimSwap(phoneNumber string, maxAge int) (*domain.CheckSimSwapInfo, error)
}

type simSwapRepository struct {
	dbURL      string
	adminUser  string
	adminPass  string
	httpClient *http.Client
}

func NewSimSwapRepository() *simSwapRepository {
	return &simSwapRepository{
		dbURL:      "http://localhost:9925",
		adminUser:  "HDB_ADMIN",
		adminPass:  "password",
		httpClient: &http.Client{},
	}
}

func (r *simSwapRepository) GetSimSwapDate(phoneNumber string) (*domain.SimSwapInfo, error) {
	payload := map[string]interface{}{
		"operation":      "search_by_hash",
		"schema":         "simswap",
		"table":          "sim_swaps",
		"hash_values":    []string{phoneNumber},
		"get_attributes": []string{"phoneNumber", "latestSimChange", "monitoredPeriod"},
	}

	responseData, err := r.executeQuery(payload)
	if err != nil {
		return nil, err
	}

	var records []map[string]interface{}
	if err := json.Unmarshal(responseData, &records); err != nil {
		return nil, fmt.Errorf("failed to parse HarperDB response: %v", err)
	}

	// If no records found, return ErrRecordNotFound
	if len(records) == 0 {
		return nil, ErrRecordNotFound
	}

	data := records[0]

	phoneNumberStr, _ := data["phoneNumber"].(string)
	latestSimChange, _ := data["latestSimChange"].(string)
	monitoredPeriodFloat, ok := data["monitoredPeriod"].(float64)
	if !ok {
		monitoredPeriodFloat = 0
	}

	return &domain.SimSwapInfo{
		PhoneNumber:     phoneNumberStr,
		LatestSimChange: latestSimChange,
		MonitoredPeriod: int(monitoredPeriodFloat),
	}, nil
}

// CheckSimSwap checks if SIM swap occurred within maxAge
func (r *simSwapRepository) CheckSimSwap(phoneNumber string, maxAge int) (*domain.CheckSimSwapInfo, error) {
	swapInfo, err := r.GetSimSwapDate(phoneNumber)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	// Compare monitored period with maxAge
	swapped := swapInfo.MonitoredPeriod < maxAge

	return &domain.CheckSimSwapInfo{
		Swapped: swapped,
	}, nil
}

// executeQuery sends a request to HarperDB
func (r *simSwapRepository) executeQuery(payload map[string]interface{}) ([]byte, error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %v", err)
	}

	req, err := http.NewRequest("POST", r.dbURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(r.adminUser, r.adminPass)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HarperDB error: %s", resp.Status)
	}

	responseData := new(bytes.Buffer)
	_, err = responseData.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	return responseData.Bytes(), nil
}
