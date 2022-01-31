#!/bin/bash
if [ $KAFKA_HOME = "" ]; then
	echo "Error: KAFKA_HOME is not set"
	exit 1
fi

if [ $# = 0 ]; then
    ${KAFKA_HOME}/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
else
    ${KAFKA_HOME}/bin/kafka-topics.sh --describe --topic $1 --bootstrap-server localhost:9092
fi
