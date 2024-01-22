package e2e

import (
	"core-goapps-orbit/pkg/customerror"
	"fmt"
	"os"
	"testing"
)

const JSONPath = "../assets/json/error_definition.json"

func TestClientExperience(t *testing.T) {
	// Read the JSON file
	jsonFile := JSONPath
	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("Error reading JSON file:%v ", err)
		return
	}

	// Load errors from JSON data
	errorStore, err := customerror.NewErrorStore(jsonData)
	if err != nil {
		t.Fatalf("Error loading errors: %v", err)
		return
	}

	// Retrieve an error message in different languages
	fmt.Println(errorStore.GetMessage("internal_error1", "en")) // English
	fmt.Println(errorStore.GetMessage("internal_error1", "es")) // Spanish
	fmt.Println(errorStore.GetMessage("internal_error1", "fr")) // French

}
func TestReadWithOutError(t *testing.T) {
	// Read the JSON file
	jsonFile := JSONPath
	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	// Load errors from JSON data
	errorStore, err := customerror.NewErrorStore(jsonData)
	if err != nil {
		t.Fatalf("Error loading errors: %v", err)
	}

	// Define test cases for different languages
	testCases := []struct {
		Code     string
		Language string
		Expected string
	}{
		{"internal_error1", "en", "Internal server error."},
		{"internal_error1", "es", "Error interno del servidor."},
		{"internal_error1", "fr", "Erreur interne du serveur."},
	}

	for _, tc := range testCases {
		msg := errorStore.GetMessage(tc.Code, tc.Language)
		if msg != tc.Expected {
			t.Errorf("GetMessage(%q, %q) = %q, want %q", tc.Code, tc.Language, msg, tc.Expected)
		}
	}
}
