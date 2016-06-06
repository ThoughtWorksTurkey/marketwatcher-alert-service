package main

import (
	"log"
	"os"
	"strings"

	"github.com/hailocab/gocassa"
)

var alertTable gocassa.Table
var cassandraClusterNodes = strings.Split(os.Getenv("CASSANDRA_NODES"), ",")
var cassandraKeySpaceName = "marketwatcher"

func getAlertTable() gocassa.Table {
	if alertTable != nil {
		return alertTable
	}

	keySpace, err := gocassa.ConnectToKeySpace(cassandraKeySpaceName, cassandraClusterNodes, "", "")

	if err != nil {
		log.Panicf("Error while connecting to Cassandra cluster: %s", err)
	}

	alertTable = keySpace.Table(
		"alert",
		&Alert{},
		gocassa.Keys{PartitionKeys: []string{"id"}},
	)

	return alertTable
}

var find = func(id int) (Alert, error) {
	result := Alert{}

	if err := getAlertTable().Where(gocassa.Eq("id", id)).ReadOne(&result).Run(); err != nil {
		return Alert{}, err
	}

	return result, nil
}

var findByOwner = func(ownerID int) ([]Alert, error) {
	results := []Alert{}
	if err := getAlertTable().Where(gocassa.Eq("owner_id", ownerID)).Read(&results).Run(); err != nil {
		return []Alert{}, err
	}

	return results, nil
}

// Upsert updates or inserts to cassandra
var upsert = func(a Alert) (Alert, error) {
	if err := getAlertTable().Set(a).Run(); err != nil {
		return Alert{}, err
	}

	return find(a.ID)
}
