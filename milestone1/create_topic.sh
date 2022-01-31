#!/bin/bash
if [ $KAFKA_HOME = "" ]; then
	echo "Error: KAFKA_HOME is not set"
	exit 1
fi

${KAFKA_HOME}/bin/kafka-topics.sh --create --topic $1 --if-not-exists --bootstrap-server localhost:9092
