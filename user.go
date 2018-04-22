package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
	"github.com/suzuki-shunsuke/go-set"
)

// User represents a user.
type User struct {
	// a unique user name used to log in with.
	// ex. "local:admin"
	Username string `json:"username,omitempty" v-create:"required" v-update:"required"`
	// the contact email address
	Email       string      `json:"email,omitempty" v-create:"required"`
	Permissions *set.StrSet `json:"permissions,omitempty" v-create:"required"`
	// a descriptive name for this account, e.g. the full name.
	FullName string `json:"full_name,omitempty" v-create:"required"`
	Password string `json:"password,omitempty" v-create:"required"`

	ID          string       `json:"id,omitempty"`
	Preferences *Preferences `json:"preferences,omitempty"`
	// the timezone to use to display times, or leave it as it is to use the system's default.
	// ex. "UTC"
	Timezone string `json:"timezone,omitempty"`
	// Session automatically end after this amount of time, unless they are actively used.
	// ex. 28800000
	SessionTimeoutMs int        `json:"session_timeout_ms,omitempty"`
	External         bool       `json:"external,omitempty"`
	Startpage        *Startpage `json:"startpage,omitempty"`
	// Assign the relevant roles to this user to grant them access to the relevant streams and dashboards.
	// The Reader role grants basic access to the system and will be enabled.
	// The Admin role grants access to everything in Graylog.
	// ex. ["Admin"]
	Roles         *set.StrSet `json:"roles,omitempty"`
	ReadOnly      bool        `json:"read_only,omitempty"`
	SessionActive bool        `json:"session_active,omitempty"`
	// ex. "2018-03-02T06:32:01.841+0000"
	LastActivity string `json:"last_activity,omitempty"`
	// ex. "192.168.192.1"
	ClientAddress string `json:"client_address,omitempty"`
}

// UserUpdateParams represents a user update API's parameter.
type UserUpdateParams struct {
	Username         string      `json:"username,omitempty" v-update:"required"`
	Email            *string     `json:"email,omitempty"`
	FullName         *string     `json:"full_name,omitempty"`
	Password         *string     `json:"password,omitempty"`
	Timezone         *string     `json:"timezone,omitempty"`
	SessionTimeoutMs *int        `json:"session_timeout_ms,omitempty"`
	Permissions      *set.StrSet `json:"permissions,omitempty"`
	Startpage        *Startpage  `json:"startpage,omitempty"`
	Roles            *set.StrSet `json:"roles,omitempty"`
}

func (user *User) NewUpdateParams() *UserUpdateParams {
	return &UserUpdateParams{
		Username:         user.Username,
		Email:            ptr.PStr(user.Email),
		FullName:         ptr.PStr(user.FullName),
		Password:         nil,
		Timezone:         ptr.PStr(user.Timezone),
		SessionTimeoutMs: ptr.PInt(user.SessionTimeoutMs),
		Permissions:      user.Permissions,
		Startpage:        user.Startpage,
		Roles:            user.Roles,
	}
}

// Preferences represents user's preferences.
type Preferences struct {
	UpdateUnfocussed  bool `json:"updateUnfocussed,omitempty"`
	EnableSmartSearch bool `json:"enableSmartSearch,omitempty"`
}

// Startpage represents a user's startpage.
type Startpage struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type UsersBody struct {
	Users []User `json:"users"`
}

// SetDefaultValues sets default values.
func (user *User) SetDefaultValues() {
	if user.SessionTimeoutMs == 0 {
		user.SessionTimeoutMs = 28800000
	}
	if user.Timezone == "" {
		user.Timezone = "UTC"
	}
}
