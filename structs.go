package alldebrid

import (
	"encoding/json"
	"time"
)

// LoginResponse ...
type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	User    User   `json:"user"`
}

type timeFromInt struct {
	time.Time
}

func (obj *timeFromInt) UnmarshalJSON(data []byte) error {
	var x int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	t := time.Unix(x, 0)
	obj.Time = t
	return nil
}

// User ...
type User struct {
	Username             string      `json:"username"`
	Email                string      `json:"email"`
	IsPremium            bool        `json:"isPremium"`
	PremiumUntil         timeFromInt `json:"premiumUntil"`
	Lang                 string      `json:"lang"`
	PreferedDomain       string      `json:"preferedDomain"`
	LimitedHostersQuotas Quotas      `json:"limitedHostersQuotas"`
}

// Quotas ...
type Quotas struct {
	SomeHost  int64 `json:"someHost"`
	OtherHost int64 `json:"otherHost"`
	OneLast   int64 `json:"oneLast"`
}

// LinkUnlockResponse ...
type LinkUnlockResponse struct {
	Success bool  `json:"success"`
	Infos   Infos `json:"infos"`
}

// Infos ...
type Infos struct {
	Link      string      `json:"link"`
	Host      string      `json:"host"`
	Filename  string      `json:"filename"`
	Streaming interface{} `json:"streaming"`
	Paws      bool        `json:"paws"`
	Filesize  string      `json:"filesize"`
}
