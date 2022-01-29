#!/bin/bash

# Create All Topics
./create_topic.sh Order
./describe_topic.sh Order 

./create_topic.sh OrderConfirmed
./describe_topic.sh OrderConfirmed

./create_topic.sh OrderPacked
./describe_topic.sh OrderPacked

./create_topic.sh Notification
./describe_topic.sh Notification 

./create_topic.sh Error
./describe_topic.sh Error

echo "modify retention of topic Orders"
./modify_retention.sh Order 259200000
./describe_topic.sh Order 

# clean up
#for topic in "Order OrderConfirmed OrderPacked Notification Error"; do ./delete_topic.sh $topic; done
