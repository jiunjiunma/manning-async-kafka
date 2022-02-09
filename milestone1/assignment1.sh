#!/bin/bash

# Create All Topics
./create_topic.sh OrderReceived
./describe_topic.sh OrderReceived

./create_topic.sh OrderConfirmed
./describe_topic.sh OrderConfirmed

./create_topic.sh OrderPacked
./describe_topic.sh OrderPacked

./create_topic.sh Notification
./describe_topic.sh Notification 

./create_topic.sh Error
./describe_topic.sh Error

echo "modify retention of topic OrderReceived"
./modify_retention.sh OrderReceived 259200000
./describe_topic.sh OrderReceived 

# clean up
#for topic in "OrderReceived OrderConfirmed OrderPacked Notification Error"; do ./delete_topic.sh $topic; done
