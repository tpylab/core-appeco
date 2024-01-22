package e2e

import (
	"core-goapps-orbit/pkg/customerror"
	"fmt"
	"os"
	"testing"
)

func TestReadCustomError(t *testing.T) {
	// Read the JSON file
	jsonFile := "../assets/json/error_definition.json"
	jsonData, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Load errors from JSON data
	errorStore, err := customerror.NewErrorStore(jsonData)
	if err != nil {
		fmt.Println("Error loading errors:", err)
		return
	}

	// Retrieve an error message in different languages
	fmt.Println(errorStore.GetMessage("internal_error1", "en")) // English
	fmt.Println(errorStore.GetMessage("internal_error1", "es")) // Spanish
	fmt.Println(errorStore.GetMessage("internal_error1", "fr")) // French

}
