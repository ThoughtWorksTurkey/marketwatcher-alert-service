package main

// Repository handles database operations
type Repository interface {
	findByOwnerID(int) ([]Alert, error)
	findByName(int, string) (Alert, error)
	upsert(Alert) (Alert, error)
}
