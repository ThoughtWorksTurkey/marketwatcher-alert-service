CREATE KEYSPACE marketwatcher WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE marketwatcher;

CREATE TABLE alert__id__ (
  id int PRIMARY KEY,
  name varchar,
  required_criteria text,
  nice_to_have_criteria text,
  excluded_criteria text,
  owner_id int,
  status int,
  threshold int
);

CREATE INDEX by_owner_id ON alert__id__(owner_id);
