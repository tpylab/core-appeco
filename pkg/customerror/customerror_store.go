package customerror

import (
	"encoding/json"
	"sync"
)

// ErrorStore stores the error messages.
type ErrorStore struct {
	sync.RWMutex
	Errors map[string]ErrorMessage
}

// NewErrorStore creates a new instance of ErrorStore.
func NewErrorStore(jsonData []byte) (*ErrorStore, error) {
	var data struct {
		Errors []ErrorEntry `json:"errors"` // Slice of error entries
	}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	em := &ErrorStore{
		Errors: make(map[string]ErrorMessage),
	}

	// Populate the error map with data from JSON
	for _, entry := range data.Errors {
		em.Errors[entry.ID] = entry.Messages
	}

	return em, nil
}

// GetMessage returns the error message in the specified language.
func (em *ErrorStore) GetMessage(errorCode, language string) string {
	em.RLock()
	defer em.RUnlock()

	// Retrieve the error message based on the error code and language
	if msg, exists := em.Errors[errorCode]; exists {
		switch language {
		case "es":
			return msg.ES
		case "en":
			return msg.EN
		case "fr":
			return msg.FR
		}
	}
	return "Unknown error"
}
