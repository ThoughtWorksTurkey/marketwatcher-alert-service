package main

import (
	"github.com/gocql/gocql"
	"log"
	"os"
	"strings"
	"testing"
)

func setup() {
	log.Println("SETTING UP")

	_, testSession := createTestConnection()
	session = testSession
	connectionEstablished = true
}

func tearDown() {
	log.Println("TEARING DOWN")

	destroyTestConnection(session)
}

func TestMain(m *testing.M) {
	setup()

	retCode := m.Run()

	tearDown()

	os.Exit(retCode)
}

var test_keyspace = "test_keyspace"

func testInitialQuery() string {
	return `
    CREATE KEYSPACE IF NOT EXISTS ` + test_keyspace + ` WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
    CREATE TABLE IF NOT EXISTS ` + test_keyspace + `.alert (
      id timeuuid PRIMARY KEY,
      name varchar,
      required_criteria text,
      nice_to_have_criteria text,
      excluded_criteria text,
      owner_id int,
      status int,
      threshold int
    );
    CREATE INDEX IF NOT EXISTS by_owner_id ON ` + test_keyspace + `.alert("owner_id");`
}

func createTestConnection() (*gocql.ClusterConfig, *gocql.Session) {
	cassandraNodes := "localhost:9042"
	cluster := gocql.NewCluster(strings.Split(cassandraNodes, ",")...)

	cluster.Timeout = cassandraConnectTimeout
	cluster.DisableInitialHostLookup = true

	var sessionErr error
	initSession, sessionErr := cluster.CreateSession()

	if sessionErr != nil {
		log.Fatal("Could NOT create test Cassandra session: " + sessionErr.Error())
	}

	executeInitialQuery(initSession, testInitialQuery())
	initSession.Close()

	cluster.Keyspace = test_keyspace
	session, _ := cluster.CreateSession()

	return cluster, session
}

func destroyTestConnection(session *gocql.Session) {
	defer session.Close()

	session.Query("truncate " + test_keyspace + ".alert").Exec()
	session.Query("drop keyspace " + test_keyspace).Exec()
}
