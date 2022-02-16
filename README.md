# manning-async-kafka
Repo used for Manning's Async Event Handling Using Kafka Live Project

## Milestone 1
Files submitted for milestone1 are located under milestone1 directory. A few wrapper scripts are created to make calling Kafka command easier. To run it, set the $KAFKA_HOME to the kafka installation location.

The 'Assignment1.sh' script contains all the steps to create the topics.

## Milestone 2
Files submitted for milestone2 are located under the milestone2 dir.

The 'service' dir contains a simple web service with only a dummy health check.

The 'producer' dir contains a kafka producer to send 3 orders to the 'OrderReceived' topic.

## Milestone 3
Files submitted for milestone3 are located under the milestone3 dir.

## Milestone 4
Files submitted for milestone4 are located under milestone4 dir. It depends on the Kafka topics created in milestone1.

### Build
```aidl
go build ./cmd/order_service
go build ./cmd/inventory_consumer
```

### Run Services
After successfully build, you will see the binary 'order_service' and 'inventory_consumer' created.
The 'order_service' is the same simple HTTP service in milestone3, which send order event to the OrderReceived topic. 
The inventory_consumer is the kafka consumer which checks for duplication and dispatches duplicated orders to the 
Error topic.

```aidl
./order_service
./inventory_consumer
```
### Test
Modify the 'send_cmd.sh' to send a test order to the order service.


