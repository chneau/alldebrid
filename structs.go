package alldebrid

import (
	"encoding/json"
	"time"
)

// LoginResponse ...
type LoginResponse struct {
	Success bool    `json:"success,omitempty"`
	Token   string  `json:"token,omitempty"`
	User    User    `json:"user,omitempty"`
	Error   *string `json:"error,omitempty"`
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
	Username             string      `json:"username,omitempty"`
	Email                string      `json:"email,omitempty"`
	IsPremium            bool        `json:"isPremium,omitempty"`
	PremiumUntil         timeFromInt `json:"premiumUntil,omitempty"`
	Lang                 string      `json:"lang,omitempty"`
	PreferedDomain       string      `json:"preferedDomain,omitempty"`
	LimitedHostersQuotas Quotas      `json:"limitedHostersQuotas,omitempty"`
}

// Quotas ...
type Quotas struct {
	SomeHost  int64 `json:"someHost,omitempty"`
	OtherHost int64 `json:"otherHost,omitempty"`
	OneLast   int64 `json:"oneLast,omitempty"`
}

// LinkUnlockResponse ...
type LinkUnlockResponse struct {
	Success bool  `json:"success,omitempty"`
	Infos   Infos `json:"infos,omitempty"`
}

// Infos ...
type Infos struct {
	Link      string      `json:"link,omitempty"`
	Host      string      `json:"host,omitempty"`
	Filename  string      `json:"filename,omitempty"`
	Streaming interface{} `json:"streaming,omitempty"`
	Paws      bool        `json:"paws,omitempty"`
	// Filesize  int         `json:"filesize"` // TODO: needs investigation
}
