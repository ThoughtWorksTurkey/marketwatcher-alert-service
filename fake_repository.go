package main

func FindByOwnerIDSuccess(id int) ([]Alert, error) {
	return []Alert{sampleAlert}, nil
}

func UpsertAlertSuccess(a Alert) (Alert, error) {
	return a, nil
}

func FindByIDSuccess(id int) (Alert, error) {
	return sampleAlert, nil
}
