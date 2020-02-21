# rabbitmq

https://www.rabbitmq.com/

RabbitMQ is a message broker, it accepts and forwards messages (binary blobs of data)
RabbitMQ speaks multiple protocols, AMQP is one of them.

A core idea of RabbitMQ is that a producer never sends any messages directly to a queue.
Instead the producer can only send messages to an exchange.

## jargon

* Producer: a program that sends messages
* Queue: messages are sent to this and stored until a consumer reads them
* Consumer: a program that receives messages
* Work Queue: distribute tasks among consumers. Round-robin by default, regardless of how many messages are queued to a consumer
* Message Acknowledgements: sent back by the consumer to tell RabbitMq that a particular message has been received, processed and RabbitMQ can delete it.
  Turned on by default
* Durable: a queue can be durable or not. If it's not durable all messages in the queue are lost if the broker restarts.
* Persistent: a message can be persistent so it's not lost if the message broker is restarted.
* Fair dispatch: by default broker sends messages to all  consumers in round robin, regardless of how busy they are. We can configure rabbitmq to not give more than one message to a consumer at any one time.
* pub/sub: deliver a message to multiple consumers
* Exchange: receives messages and pushes them to queue(s)
  It knows exactly what to do with the message based on the queue type.
  4 types of exchanges:
    * direct
      a message goes to the queues whose binding key exactly matches the routing key of the message
      multiple queues can have the same key
      and multiple keys can point to one queue
    * topic
      a message goes to the queues whose bindings have a wildcard match to the key
    * fanout
      broadcasts all messages it receives to all queues it knows about
    * headers
      use message header attributes for routing
  By default there is a nameless exchange which is used if you send a message directly to a queue.
* Routing: subscribe a consumer to only a subset of messages coming from an exchange

## implementation

Use a *ConnectionFactory* to configure a connection to a rabbitmq message broker.
Create a *Connection* from the factory.
Create a *Channel* from the Connection.
  Most of the APIs for getting things done reside
Declare a *Queue* from the channel.
  Declaring a queue is idempotent, i.e. if it already exists it will be reused, otherwise created.
The message contents is a byte array
A *DeliverCallback* is a callback which is executed when a message is received.

### exchanges

Declare an exchange on the *Channel*
publish to this exchange.
Declare a queue on the *Channel*
  since we don't need the name we can let Rabbit generate it
On the *Channel* bind the queue to the exchange.
  Once bound any messages sent to the exchange will be forwarded to the queue
