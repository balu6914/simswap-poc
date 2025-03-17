package usecases

type SimSwapUsecase interface {
	RetrieveSimSwapDate(phoneNumber string) (string, error)
	CheckSimSwap(phoneNumber string, maxAge int) (bool, error)
}
