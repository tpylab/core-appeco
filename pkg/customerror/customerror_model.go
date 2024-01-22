package customerror

// ErrorMessage stores error messages in different languages.
type ErrorMessage struct {
	ES string `json:"es"` // Spanish message
	EN string `json:"en"` // English message
	FR string `json:"fr"` // French message
}

// ErrorEntry represents an individual error entry.
type ErrorEntry struct {
	ID       string       `json:"id"`
	Type     string       `json:"type"`
	Messages ErrorMessage `json:"messages"`
}
