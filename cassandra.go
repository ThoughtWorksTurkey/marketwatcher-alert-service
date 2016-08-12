package main

import (
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

var save = func(a Alert) (Alert, error) {
	if connectionEstablished == false {
		connectToCassandra()
		connectionEstablished = true
	}

	if err := session.Query(`INSERT INTO alert (id, owner_id, name, required_criteria, nice_to_have_criteria, excluded_criteria, threshold, status)
	VALUES (?,?,?,?,?,?,?,?)`, a.ID, a.OwnerID, a.Name, a.RequiredCriteria, a.NiceToHaveCriteria, a.ExcludedCriteria, a.Threshold, a.Status).Exec(); err != nil {
		log.Fatal(err)
	}

	return find(a.ID)
}

func connectToCassandra() {
	cluster := gocql.NewCluster(strings.Split(os.Getenv("CASSANDRA_NODES"), ",")...)
	cluster.Timeout = cassandraConnectTimeout

	executeInitialQuery(cluster)

	cluster.Keyspace = cassandraKeySpaceName

	var sessionErr error
	session, sessionErr = cluster.CreateSession()
	if sessionErr != nil {
		log.Fatal("Could NOT create Cassandra session: ", sessionErr)
	}
}

func executeInitialQuery(cl *gocql.ClusterConfig) {
	initialBytes, err := ioutil.ReadFile(initialCqlFile)
	initialQuery := string(initialBytes)

	initSession, err := cl.CreateSession()
	defer initSession.Close()
	if err != nil {
		log.Fatal("Could NOT create a session for initial CQL file: ", err)
	}

	for _, l := range strings.Split(initialQuery, ";") {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		log.Println("Executing Initial query: ", l)

		if err = initSession.Query(l).Exec(); err != nil {
			log.Fatal("Could NOT execute initial query: ", err)
		}
	}
}
