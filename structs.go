package alldebrid

// LoginResponse ...
type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	User    User   `json:"user"`
}

// User ...
type User struct {
	Username             string `json:"username"`
	Email                string `json:"email"`
	IsPremium            bool   `json:"isPremium"`
	PremiumUntil         int    `json:"premiumUntil"`
	Lang                 string `json:"lang"`
	PreferedDomain       string `json:"preferedDomain"`
	LimitedHostersQuotas Quotas `json:"limitedHostersQuotas"`
}

// Quotas ...
type Quotas struct {
	SomeHost  int64 `json:"someHost"`
	OtherHost int64 `json:"otherHost"`
	OneLast   int64 `json:"oneLast"`
}
