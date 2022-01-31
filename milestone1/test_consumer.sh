#!/bin/bash
if [ $KAFKA_HOME = "" ]; then
	echo "Error: KAFKA_HOME is not set"
	exit 1
fi

${KAFKA_HOME}/bin/kafka-console-consumer.sh --topic $1 --from-beginning --bootstrap-server localhost:9092
