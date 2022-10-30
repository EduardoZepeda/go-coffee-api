package models

type Feed struct {
	Username string `db:"username" json:"user,omitempty"`
	Action   string `db:"action" json:"verb,omitempty"`
	Target   string `db:"target" json:"object,omitempty"`
}
