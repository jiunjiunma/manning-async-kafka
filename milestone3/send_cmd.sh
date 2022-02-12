#!/bin/bash

curl http://localhost:8080/orders \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","name": "order4", "customer_id": "customer1", "items":[ {"quantity":1, "item_id": "itemOne"}]}'
