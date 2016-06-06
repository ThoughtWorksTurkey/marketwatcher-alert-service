package main

import (
	"github.com/hailocab/gocassa"
	"log"
)

var alertTable gocassa.Table
var cassandraClusterNodes = []string{"192.168.1.200"}
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


var find = func (id int) (Alert, error) {
	result := Alert{}

	if err := getAlertTable().Where(gocassa.Eq("id", id)).ReadOne(&result).Run(); err != nil {
		return Alert{}, err
	}

	return result, nil
}

// Upsert updates or inserts to cassandra
var upsert = func (a Alert) (Alert, error) {
	if err := getAlertTable().Set(a).Run(); err != nil {
		return Alert{}, err
	}

	return find(a.ID)
}
