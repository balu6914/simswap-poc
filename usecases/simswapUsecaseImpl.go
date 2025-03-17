package usecases

import "simswap-poc/repositories"

type SimSwapUsecaseImpl struct {
	Repo repositories.SimSwapRepository
}

func NewSimSwapUsecase(repo repositories.SimSwapRepository) *SimSwapUsecaseImpl {
	return &SimSwapUsecaseImpl{Repo: repo}
}

func (u *SimSwapUsecaseImpl) RetrieveSimSwapDate(phoneNumber string) (string, error) {
	return u.Repo.GetLatestSimSwapDate(phoneNumber)
}

func (u *SimSwapUsecaseImpl) CheckSimSwap(phoneNumber string, maxAge int) (bool, error) {
	return u.Repo.CheckSimSwapInPeriod(phoneNumber, maxAge)
}
