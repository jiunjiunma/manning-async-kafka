#!/bin/bash
if [ $KAFKA_HOME = "" ]; then
	echo "Error: KAFKA_HOME is not set"
	exit 1
fi

${KAFKA_HOME}/bin/kafka-console-producer.sh --topic $1 --bootstrap-server localhost:9092
