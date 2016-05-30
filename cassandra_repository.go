package main

import (
	"github.com/gocql/gocql"
	"log"
)

var session *gocql.Session

// CassandraRepository is a fake repository
type CassandraRepository struct{}

func initializeDB() {
	cluster := gocql.NewCluster("192.168.1.200")
	cluster.ProtoVersion = 4
	cluster.Keyspace = "marketwatcher"
	ssn, _ := cluster.CreateSession()
	session = ssn
}

func (cr CassandraRepository) find(id int) (Alert, error) {
	if session == nil {
		initializeDB()
	}
	a := Alert{}
	if err := session.Query("SELECT * FROM alert WHERE id=?", id).Scan(&a); err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func (cr CassandraRepository) upsert(a Alert) (Alert, error) {
	if session == nil {
		initializeDB()
	}
	return Alert{}, nil
}
