package types

type ApiError struct {
	Message string             `json:"message"`
	Errors  *map[string]string `json:"errors,omitempty"`
}
