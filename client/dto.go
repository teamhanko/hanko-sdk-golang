package client

type OperationStatus string

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IconUrl     string `json:"icon"`
}

type ApiError struct {
	Message      string `json:"message,omitempty"`
	Details      string `json:"details,omitempty"`
	DebugMessage string `json:"debug_message,omitempty"`
	StatusText   string `json:"status_text"`
	StatusCode   int    `json:"status_code"`
}

