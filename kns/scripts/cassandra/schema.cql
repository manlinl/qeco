CREATE KEYSPACE kns
    WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

CREATE TABLE kns.endpoints
(
    kns_name text,
    address  blob,
    labels   map<text,text>,
    PRIMARY KEY (kns_name, address)
);