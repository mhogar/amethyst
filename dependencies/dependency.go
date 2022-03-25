package dependencies

import "sync"

type DependencyFactory[T any] func() T

type Dependency[T any] struct {
	once         sync.Once
	object       T
	createObject DependencyFactory[T]
}

// Resolve resolves the dependency.
// Only the first call to this function will create a new object, after which it will be retrieved from memory.
func (d Dependency[T]) Resolve() T {
	d.once.Do(func() {
		d.object = d.createObject()
	})
	return d.object
}
