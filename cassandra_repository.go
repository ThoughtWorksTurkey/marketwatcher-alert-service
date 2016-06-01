package main

import (
	"github.com/hailocab/gocassa"
	"log"
)

var alertTable gocassa.Table
var cassandraClusterNodes = []string{"192.168.1.200"}
var cassandraKeySpaceName = "marketwatcher"

// CassandraRepository is a fake repository
type CassandraRepository struct{}

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

func (cr CassandraRepository) find(id int) (Alert, error) {
	result := Alert{}

	if err := getAlertTable().Where(gocassa.Eq("id", id)).ReadOne(&result).Run(); err != nil {
		return Alert{}, err
	}

	return result, nil
}

func (cr CassandraRepository) upsert(a Alert) (Alert, error) {

	return Alert{}, nil
}
