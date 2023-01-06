#!/bin/bash
[ -d /tmp/cassandra ] || mkdir -p /tmp/cassandra
UserID="$(id -u)" GroupID="$(id -g)" docker-compose up -d
