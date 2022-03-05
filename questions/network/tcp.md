# TCP面试题整理

## 简述 TCP 三次握手以及四次挥手的流程。为什么需要三次握手以及四次挥手？

三次握手：

- 连接的发起方发送SYN和SEQ，进入SYN_SENT状态。
- 接收方收到之后发送SYN_ACK和SEQ，进入SYN_RECVED状态。
- 发起方发送ACK，双方连接建立。

因为TCP连接是可靠的。所以在建立连接的时候，必须确保双方都可以正常接收和发送数据。第一次握手是客户端确认服务端可以接受数据；第二次客户端确认了服务端可以发送数据，并且服务端确认了客户端可以接受数据；最后一次，服务端确认了客户端可以发送数据，二者连接建立。

除了确保双方正常发送接受数据，TCP还需要确保数据包是可以重发的，因此需要双方交换序列号SEQ，以防止数据被重复接收。本来SEQ需要在确认连接之后再单独交换，但是TCP协议复用了这一过程，直接在确认连接的过程中把SEQ的交换也给完成了。所以我们不需要4次握手而是3次握手。

四次挥手：

- 当客户端没有数据要发送的时候，会发送一个FIN，并进入FIN_WAIT_1状态。
- 服务端确认客户端没有数据要发送，返回FIN_ACK，进入FIN_WAIT_2状态。
- 这时服务端还是可以正常给客户端发送数据。而当服务端也没有数据要发送了，就发送FIN给客户端。
- 客户端收到FIN，返回FIN_ACK关闭连接。

因为TCP是全双工协议，双方都可以给对方发送数据，在一方没有数据要发送时，另外一方可能还有数据没有发送完毕。因此发送方和接收方各需要两次挥手来确认没有数据发送了。

## TCP 怎么保证可靠传输？

- 每个数据包发送之后，必须等到对方确认收到了，才会认为这个数据包已经成功发送了。
- 当超时没有收到数据包的确认信息，认为这个数据包丢了，会重新发送。通过SEQ来确保数据包不会因为重传导致重复接收。
- 使用滑动窗口进行流量控制，每个窗口控制数据包的发送量，防止因为发送过量数据包导致对方难以处理。
- 拥塞控制，当丢包过多时，启用拥塞控制算法，减少数据包的发送频率。

## TIME_WAIT是什么？它的作用是什么？

TIME_WAIT出现在主动关闭连接的那一方。因为在TCP连接关闭之后，可能会出现延时数据包的情况，对方可能还有数据包残留在网络中没有被接收。这时候连接关闭方会等待一段时间（一般是2WSL）以确保所有数据都被接收了。另一个作用是，等待一段时间确保FIN_ACK被对面接收到，以正确关闭连接。

## TCP的KeepAlive是什么？作用是什么？

在TCP长连接中，如果长期没有收到对方的数据，可能触发timeout导致连接断开。KeepAlive在双方没有数据交互时通过心跳包来确认双方的存活。