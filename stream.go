package graylog

// CloneStream
// POST /streams/{streamID}/clone Clone a stream
// TestMatchStream
// POST /streams/{streamID}/testMatch Test matching of a stream against a supplied message

// Stream represents a steram.
type Stream struct {
	// required
	Title string `json:"title,omitempty" v-create:"required"`
	// ex. "5a8c086fc006c600013ca6f5"
	IndexSetID string `json:"index_set_id,omitempty" v-create:"required"`

	// ex. "5a94abdac006c60001f04fc1"
	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`
	// ex. "2018-02-20T11:37:19.371Z"
	CreatedAt string `json:"created_at,omitempty" v-create:"isdefault"`
	// ex. local:admin
	CreatorUserID string   `json:"creator_user_id,omitempty" v-create:"isdefault"`
	Description   string   `json:"description,omitempty"`
	Outputs       []Output `json:"outputs,omitempty" v-create:"isdefault"`
	// ex. "AND"
	MatchingType                   string           `json:"matching_type,omitempty"`
	Disabled                       bool             `json:"disabled,omitempty" v-create:"isdefault"`
	Rules                          []StreamRule     `json:"rules,omitempty"`
	AlertConditions                []AlertCondition `json:"alert_conditions,omitempty" v-create:"isdefault"`
	AlertReceivers                 *AlertReceivers  `json:"alert_receivers,omitempty" v-create:"isdefault"`
	RemoveMatchesFromDefaultStream bool             `json:"remove_matches_from_default_stream,omitempty"`
	IsDefault                      bool             `json:"is_default,omitempty" v-create:"isdefault"`
	// ContentPack `json:"content_pack,omitempty"`
}

// Output represents an output.
type Output struct{}

// AlertReceivers represents alert receivers.
type AlertReceivers struct {
	Emails []string `json:"emails,omitempty"`
	Users  []string `json:"users,omitempty"`
}

// AlertCondition represents an alert condition.
type AlertCondition struct{}

type StreamsBody struct {
	Total   int      `json:"total,omitempty"`
	Streams []Stream `json:"streams,omitempty"`
}
