package main

// Repository handles database operations
type Repository interface {
	find(int) (Alert, error)
	upsert(Alert) (Alert, error)
}
