# RabbitMQ

## 交换机有哪几种类型？

- fanout: 交换机收到的数据会转发到它绑定的所有队列中。
- direct: 把消息发送到bindkey和routekey完全匹配的队列。
- topic: bindkey可以动态匹配routekey。

## rabbitmq跟kafka对比？

RabbitMQ是一个实现了AMQP协议的消息系统，它的实时性跟可靠性很高；而kafka是一个分布式流式系统。

RabbitMQ一般用在实时业务系统中，而kafka一般用在异步高吞吐量的场景，例如离线大数据分析，异步数据收集，异步监控数据收集等。