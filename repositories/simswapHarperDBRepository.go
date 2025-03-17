package repositories

import (
	"encoding/json"
	"fmt"
	"simswap-poc/database"
	"simswap-poc/entities"
)

type SimSwapHarperDBRepository struct {
	Client *database.HarperDBClient
}

func NewSimSwapHarperDBRepository(client *database.HarperDBClient) *SimSwapHarperDBRepository {
	return &SimSwapHarperDBRepository{Client: client}
}

func (r *SimSwapHarperDBRepository) GetLatestSimSwapDate(phoneNumber string) (string, error) {
	operation := map[string]interface{}{
		"operation": "sql",
		"sql":       fmt.Sprintf("SELECT latestSimChange FROM dev.simswap WHERE phoneNumber = '%s'", phoneNumber),
	}

	resp, err := r.Client.ExecuteOperation(operation)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result []entities.SimSwap
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", nil
	}

	return result[0].LatestSimChange, nil
}

func (r *SimSwapHarperDBRepository) CheckSimSwapInPeriod(phoneNumber string, maxAge int) (bool, error) {
	operation := map[string]interface{}{
		"operation": "sql",
		"sql":       fmt.Sprintf("SELECT latestSimChange FROM dev.simswap WHERE phoneNumber = '%s'", phoneNumber),
	}

	resp, err := r.Client.ExecuteOperation(operation)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result []entities.SimSwap
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if len(result) == 0 {
		return false, nil
	}

	// Logic to check if the SIM swap occurred within the maxAge period
	// (You can implement this based on your requirements)

	return true, nil
}
