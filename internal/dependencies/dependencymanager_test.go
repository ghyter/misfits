package dependencies_test

import (
	"testing"

	"github.com/ghyter/misfits/internal/dependencies"
)

type MockDependency struct {
	Name string
}

func TestDependencyManager_RegisterAndGet(t *testing.T) {
	dm := dependencies.NewDependencyManager()

	// Register a dependency
	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (*MockDependency, error) {
		return &MockDependency{Name: "TestDependency"}, nil
	})

	// Retrieve the dependency
	instance, err := dependencies.Get[*MockDependency](dm)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if instance.Name != "TestDependency" {
		t.Errorf("Expected dependency name 'TestDependency', got '%s'", instance.Name)
	}
}

func TestDependencyManager_DuplicateRegister(t *testing.T) {
	dm := dependencies.NewDependencyManager()

	// Register a dependency
	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (*MockDependency, error) {
		return &MockDependency{Name: "FirstDependency"}, nil
	})

	// Attempt to register the same type again
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Expected panic for duplicate registration, but none occurred")
		}
	}()

	dependencies.Register(dm, func(dm *dependencies.DependencyManager) (*MockDependency, error) {
		return &MockDependency{Name: "DuplicateDependency"}, nil
	})
}
