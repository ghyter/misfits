package dependencies

import (
	"fmt"
	"reflect"
	"sync"
)

// DependencyManager manages dependencies by initializing them at registration.
type DependencyManager struct {
	instances map[reflect.Type]any // Cache for initialized dependencies
	instLock  sync.Mutex           // Lock for accessing instances
	regLock   sync.Mutex           // Lock for accessing registry
}

// NewDependencyManager creates a new instance of DependencyManager.
func NewDependencyManager() *DependencyManager {
	return &DependencyManager{
		instances: make(map[reflect.Type]any),
	}
}

// Register initializes and registers a dependency for the specified type.
func Register[T any](dm *DependencyManager, initializer func(*DependencyManager) (T, error)) {
	dm.regLock.Lock()
	defer dm.regLock.Unlock()

	typ := reflect.TypeOf((*T)(nil)).Elem()
	if _, exists := dm.instances[typ]; exists {
		panic(fmt.Sprintf("dependency of type %s is already registered", typ))
	}

	// Initialize the dependency immediately using the DependencyManager
	instance, err := initializer(dm)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize dependency of type %s: %v", typ, err))
	}

	dm.instances[typ] = instance
}

// Get retrieves the already-initialized dependency for the specified type.
func Get[T any](dm *DependencyManager) (T, error) {
	dm.instLock.Lock()
	defer dm.instLock.Unlock()

	typ := reflect.TypeOf((*T)(nil)).Elem()

	// Retrieve the cached instance
	instance, exists := dm.instances[typ]
	if !exists {
		var zero T
		return zero, fmt.Errorf("no dependency registered for type %s", typ)
	}

	return instance.(T), nil
}
