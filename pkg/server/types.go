package server

// APIResponse represents a successful JSON response containing a message and optional data.
type APIResponse struct {
	Message string `json:"message"` // The response message to be returned in JSON format
	Data    any    `json:"data"`    // Optional data to be returned in JSON format
}

// APIError represents an error response in JSON format.
type APIError struct {
	Error string `json:"error"` // The error message to be returned in JSON format
}
