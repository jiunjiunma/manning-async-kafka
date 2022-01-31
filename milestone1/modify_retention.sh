#!/bin/bash 
if [ $KAFKA_HOME = "" ]; then
	echo "Error: KAFKA_HOME is not set"
	exit 1
fi

${KAFKA_HOME}/bin/kafka-configs.sh --alter \
      --add-config retention.ms=$2 \
      --bootstrap-server=0.0.0.0:9092 \
      --topic $1
