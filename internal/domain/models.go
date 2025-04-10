package domain

type SimSwapInfo struct {
	PhoneNumber     string `json:"phoneNumber"` // ✅ Add this field
	LatestSimChange string `json:"latestSimChange"`
	MonitoredPeriod int    `json:"monitoredPeriod"`
}

type CheckSimSwapInfo struct {
	Swapped bool `json:"swapped"`
}

type CreateSimSwapDate struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CreateCheckSimSwap struct {
	PhoneNumber string `json:"phoneNumber"`
	MaxAge      int    `json:"maxAge"`
}


