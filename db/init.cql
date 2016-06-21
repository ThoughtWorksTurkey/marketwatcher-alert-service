CREATE KEYSPACE IF NOT EXISTS marketwatcher WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
CREATE TABLE IF NOT EXISTS marketwatcher.alert (
  id timeuuid PRIMARY KEY,
  name varchar,
  required_criteria text,
  nice_to_have_criteria text,
  excluded_criteria text,
  owner_id int,
  status int,
  threshold int
);
CREATE INDEX IF NOT EXISTS by_owner_id ON marketwatcher.alert("owner_id");
