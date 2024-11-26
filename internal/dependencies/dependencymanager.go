// internal/dependencies/dependency_manager.go
package dependencies

import (
	"fmt"
	"reflect"
	"sync"
)

// DependencyManager is a generic manager for lazy initialization and retrieval of dependencies.
type DependencyManager struct {
	registry  map[reflect.Type]func() (any, error) // Registry of dependency initializers
	instances map[reflect.Type]any                 // Cache for initialized dependencies
	lock      sync.Mutex                           // Mutex for thread-safe access
}

// NewDependencyManager creates a new instance of DependencyManager.
func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		registry:  make(map[reflect.Type]func() (any, error)),
		instances: make(map[reflect.Type]any),
	}
}

// Register adds a new dependency initializer for the specified type.
func Register[T any](dm *DependencyManager, initializer func() (T, error)) {
	dm.lock.Lock()
	defer dm.lock.Unlock()

	typ := reflect.TypeOf((*T)(nil)).Elem()
	if _, exists := dm.registry[typ]; exists {
		panic(fmt.Sprintf("dependency of type %s is already registered", typ))
	}

	// Wrap initializer to match the generic interface
	dm.registry[typ] = func() (any, error) {
		return initializer()
	}
}

// Get retrieves the dependency for the specified type, initializing it if necessary.
func Get[T any](dm *DependencyManager) (T, error) {
	dm.lock.Lock()
	defer dm.lock.Unlock()

	typ := reflect.TypeOf((*T)(nil)).Elem()

	// Check if already initialized
	if instance, exists := dm.instances[typ]; exists {
		return instance.(T), nil
	}

	// Initialize dependency
	initializer, exists := dm.registry[typ]
	if !exists {
		var zero T
		return zero, fmt.Errorf("no dependency registered for type %s", typ)
	}

	instance, err := initializer()
	if err != nil {
		var zero T
		return zero, fmt.Errorf("failed to initialize dependency of type %s: %w", typ, err)
	}

	// Cache the initialized instance
	dm.instances[typ] = instance
	return instance.(T), nil
}
