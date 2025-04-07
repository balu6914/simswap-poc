package usecase

import (
	"errors"
	"fmt"
	"simswap-poc/internal/domain"
	"simswap-poc/internal/repository"
)

type SimSwapUsecase interface {
	RetrieveSimSwapDate(req domain.CreateSimSwapDate) (*domain.SimSwapInfo, *domain.ErrorResponse)
	CheckSimSwap(req domain.CreateCheckSimSwap) (*domain.CheckSimSwapInfo, *domain.ErrorResponse)
}

type simSwapUsecase struct {
	repo repository.SimSwapRepository
}

func NewSimSwapUsecase(repo repository.SimSwapRepository) SimSwapUsecase {
	return &simSwapUsecase{repo: repo}
}

func (uc *simSwapUsecase) RetrieveSimSwapDate(req domain.CreateSimSwapDate) (*domain.SimSwapInfo, *domain.ErrorResponse) {
	if req.PhoneNumber == "" {
		return nil, domain.Generic422 // Validation failed
	}

	if !uc.isAuthorized(req) { // Placeholder; implement logic in the time of real data
		return nil, domain.Generic401
	}

	// Simulate forbidden access (e.g., user lacks permission)
	if !uc.hasPermission(req) { // Placeholder; implement logic in the time of real data
		return nil, domain.Generic403
	}

	info, err := uc.repo.GetSimSwapDate(req.PhoneNumber)
	fmt.Println("Repository returned error:", err)
	if errors.Is(err, repository.ErrRecordNotFound) {
		return nil, domain.Generic404
	} else if err != nil {
		return nil, domain.Generic400
	}

	// Simulate rate limiting (e.g., too many requests)
	if uc.isRateLimited(req) { // Placeholder; implement your logic
		return nil, domain.Generic429
	}

	return info, nil
}

func (uc *simSwapUsecase) CheckSimSwap(req domain.CreateCheckSimSwap) (*domain.CheckSimSwapInfo, *domain.ErrorResponse) {

	if req.PhoneNumber == "" || req.MaxAge < 0 {
		return nil, domain.Generic422
	}

	if !uc.isAuthorized(req) {
		return nil, domain.Generic401
	}

	if !uc.hasPermission(req) {
		return nil, domain.Generic403
	}

	info, err := uc.repo.CheckSimSwap(req.PhoneNumber, req.MaxAge)
	if err != nil {
		return nil, domain.Generic400
	}

	if uc.isRateLimited(req) {
		return nil, domain.Generic429
	}

	return info, nil
}

func (uc *simSwapUsecase) isAuthorized(req interface{}) bool {
	return false // Will Replace with real auth logic in the time of real time data
}

func (uc *simSwapUsecase) hasPermission(req interface{}) bool {
	return true // will Replace with real permission logic in the time of real time data ,
}

func (uc *simSwapUsecase) isRateLimited(req interface{}) bool {
	return false // will Replace with real rate-limiting logic in the time of real time data
}
