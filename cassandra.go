package main

import (
	"errors"
	"github.com/gocql/gocql"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var cassandraKeySpaceName = "marketwatcher"
var session *gocql.Session
var cassandraConnectTimeout = 10 * time.Second
var initialCqlFile = "/data/init.cql"
var connectionEstablished = false

const ALERT_NAME_MUST_BE_UNIQUE_PER_OWNER = "Alert name must be unique per owner"

var find = func(id gocql.UUID) (Alert, error) {
	if connectionEstablished == false {
		connectToCassandra()
		connectionEstablished = true
	}

	result := Alert{}

	if err := session.Query(`SELECT id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status FROM alert WHERE id = ?`,
		id).Consistency(gocql.One).Scan(&result.ID, &result.OwnerID, &result.Name, &result.RequiredCriteria, &result.NiceToHaveCriteria, &result.ExcludedCriteria, &result.Threshold, &result.Status); err != nil {
		log.Println("Find returns no records, check given parameters: ", err)
	}

	return result, nil
}

var findByOwner = func(ownerID int) ([]Alert, error) {
	if connectionEstablished == false {
		connectToCassandra()
		connectionEstablished = true
	}

	result := Alert{}
	alerts := []Alert{}
	iter := session.Query(`SELECT id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status FROM alert WHERE owner_id = ?`, strconv.Itoa(ownerID)).Iter()
	for iter.Scan(&result.ID, &result.OwnerID, &result.Name, &result.RequiredCriteria, &result.NiceToHaveCriteria, &result.ExcludedCriteria, &result.Threshold, &result.Status) {
		alerts = append(alerts, result)
	}

	err := iter.Close()
	return alerts, err
}

var alertExists = func(session *gocql.Session, ownerID int, name string) bool {
	count := 0
	session.Query(`SELECT count(*) FROM alert WHERE owner_id = ? AND name = ? ALLOW FILTERING`, strconv.Itoa(ownerID), name).Consistency(gocql.One).Scan(&count)

	return count > 0
}

var save = func(a Alert) (Alert, error) {
	if connectionEstablished == false {
		connectToCassandra()
		connectionEstablished = true
	}

	if alertExists(session, a.OwnerID, a.Name) {
		return a, errors.New(ALERT_NAME_MUST_BE_UNIQUE_PER_OWNER)
	}

	if err := session.Query(`INSERT INTO alert (id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status)
	VALUES (?,?,?,?,?,?,?,?)`, a.ID, a.OwnerID, a.Name, a.RequiredCriteria, a.NiceToHaveCriteria, a.ExcludedCriteria, a.Threshold, a.Status).Exec(); err != nil {
		return a, err
	}

	return find(a.ID)
}

var initialQuery = func() string {
	initialBytes, _ := ioutil.ReadFile(initialCqlFile)
	return string(initialBytes)
}

func connectToCassandra() error {
	cassandraNodes := os.Getenv("CASSANDRA_NODES")

	if cassandraNodes == "" {
		log.Fatal("Please specify CASSANDRA_NODES env variable")
	}

	cluster := gocql.NewCluster(strings.Split(cassandraNodes, ",")...)
	cluster.Timeout = cassandraConnectTimeout
	cluster.Keyspace = cassandraKeySpaceName
	cluster.DisableInitialHostLookup = true

	initSession, sessionErr := cluster.CreateSession()

	if sessionErr != nil {
		log.Fatal("Could NOT create Cassandra session: " + sessionErr.Error())
	}

	executeInitialQuery(initSession, initialQuery())
	initSession.Close()

	session, sessionErr = cluster.CreateSession()

	return nil
}

func executeInitialQuery(session *gocql.Session, initialQuery string) error {
	for _, l := range strings.Split(initialQuery, ";") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		if err := session.Query(l).Exec(); err != nil {
			return errors.New("Could NOT execute initial query: " + err.Error())
		}
	}
	return nil
}
