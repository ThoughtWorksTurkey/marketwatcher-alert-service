package main

import (
	"log"
	"os"
	"strings"
	"time"
	"github.com/gocql/gocql"
	"strconv"
)

var cassandraKeySpaceName = "marketwatcher"
var session *gocql.Session

func init() {
	cluster := gocql.NewCluster(strings.Split(os.Getenv("CASSANDRA_NODES"), ",")...)
	cluster.Keyspace = cassandraKeySpaceName
	cluster.Timeout = 1 * time.Second
	session, _ = cluster.CreateSession()
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
	iter := session.Query(`SELECT id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status FROM alert__id__ WHERE owner_id = ?`, strconv.Itoa(ownerID)).Iter()
	for iter.Scan(&result.ID, &result.OwnerID, &result.Name, &result.RequiredCriteria, &result.NiceToHaveCriteria, &result.ExcludedCriteria, &result.Threshold, &result.Status) {
		alerts = append(alerts, result)
	}
	err := iter.Close()
	return alerts, err
}

var save = func(a Alert) (Alert, error) {
	a.ID = GenerateAlertId()

	if err := session.Query(`INSERT INTO alert__id__ (id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status)
	VALUES (?,?,?,?,?,?,?,?)`, a.ID, a.OwnerID, a.Name, a.RequiredCriteria, a.NiceToHaveCriteria, a.ExcludedCriteria, a.Threshold, a.Status).Exec(); err != nil {
		log.Fatal(err)
	}

	return find(a.ID)
}
