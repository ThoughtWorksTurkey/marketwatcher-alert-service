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
		gocassa.Keys{PartitionKeys: []string{"owner_id", "name"}},
	)

	alertTable.CreateIfNotExist()

	return alertTable
}

func find(id int) (Alert, error) {
	result := Alert{}

	if err := getAlertTable().Where(gocassa.Eq("id", id)).ReadOne(&result).Run(); err != nil {
		return Alert{}, err
	}

	return result, nil
}

func PersistAlert(a Alert) (Alert, error) {
	return Alert{}, nil
}
