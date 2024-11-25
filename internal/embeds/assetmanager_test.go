package embeds

import (
	"testing"
)

// TestNewAssetManager verifies that a new AssetManager can be created successfully.
func TestNewAssetManager(t *testing.T) {
	manager := NewDefaultAssetManager()
	if manager == nil {
		t.Fatalf("Expected AssetManager instance, got nil")
	}
}

// TestGet_ExistingFile verifies that the AssetManager can retrieve an embedded file successfully.
func TestGet_ExistingFile(t *testing.T) {
	manager := NewDefaultAssetManager()

	// Use a file known to exist in the embedded assets
	data, err := manager.Get("fonts/DejaVuSans.ttf")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Verify the content is non-empty
	if len(data) == 0 {
		t.Fatalf("Expected non-empty data for existing file")
	}
}

// TestGet_NonExistentFile verifies that the AssetManager returns an error for a missing file.
func TestGet_NonExistentFile(t *testing.T) {
	manager := NewDefaultAssetManager()

	// Try to get a file that doesn't exist
	_, err := manager.Get("nonexistent-file.txt")
	if err == nil {
		t.Fatalf("Expected error for missing file, got nil")
	}

	// Optional: Check the error message contains the missing filename
	expectedError := "failed to load asset"
	if !contains(err.Error(), expectedError) {
		t.Fatalf("Expected error containing '%s', got: %v", expectedError, err)
	}
}

// contains checks if a substring is present in a string.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr || len(s) > len(substr) && contains(s[1:], substr)
}
