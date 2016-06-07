package main

import (
	"log"
	"os"
	"strings"

	"github.com/hailocab/gocassa"
	"github.com/gocql/gocql"
)

var alertTable gocassa.Table
var cassandraClusterNodes = strings.Split(os.Getenv("CASSANDRA_NODES"), ",")
var cassandraKeySpaceName = "marketwatcher"
var session *gocql.Session

func init() {
	cluster := gocql.NewCluster("192.168.1.200")
	cluster.Keyspace = cassandraKeySpaceName
	session, _ = cluster.CreateSession()
}

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

var find = func(id gocql.UUID) (Alert, error) {
	result := Alert{}

	if err := session.Query(`SELECT id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status FROM alert__id__ WHERE id = ?`,
		id).Consistency(gocql.One).Scan(&result.ID, &result.OwnerID, &result.Name, &result.RequiredCriteria, &result.NiceToHaveCriteria, &result.ExcludedCriteria, &result.Threshold, &result.Status); err != nil {
		log.Fatal(err)
	}

	return result, nil
}

var findByOwner = func(ownerID int) ([]Alert, error) {
	result := Alert{}
	alerts := []Alert{}
	iter := session.Query(`SELECT id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status FROM alert__id__ WHERE owner_id = 1`).Iter()
	for iter.Scan(&result.ID, &result.OwnerID, &result.Name, &result.RequiredCriteria, &result.NiceToHaveCriteria, &result.ExcludedCriteria, &result.Threshold, &result.Status) {
		alerts = append(alerts, result);
	}
	err := iter.Close()
	return alerts, err
}

// Upsert updates or inserts to cassandra
var save = func(a Alert) (Alert, error) {
	a.ID = GenerateAlertId()
	if err := getAlertTable().Set(a).Run(); err != nil {
		return Alert{}, err
	}

	return find(a.ID)
}
