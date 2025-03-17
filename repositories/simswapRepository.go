package repositories

type SimSwapRepository interface {
	GetLatestSimSwapDate(phoneNumber string) (string, error)
	CheckSimSwapInPeriod(phoneNumber string, maxAge int) (bool, error)
}
