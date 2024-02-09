# counter service


### 
1. 写性能差 --> 异步写
    1.1 在计数服务端聚合写  (缺点: 中途宕机会丢数据，少记)
        适合少记不敏感的数据 (例如: 浏览量) 

    1.2 引入消息队列，在消费者端聚合写  (缺点: 中途宕机会重复消费，多记)
        适合少记敏感的数据 (例如: 点赞、收藏)
        At least one 原则, 要在消费端采用 布隆过滤器+唯一索引 的方案，实现业务幂等性的同时支撑业务高并发

2. 读性能差 --> 缓存数据
    
      1. read through + write through ?
      2. singlefight 

    
### 
1. define ddl sql
2. goctl model mysql ddl --src model/ddl.sql --dir .
3. define .proto file
4. goctl rpc protoc go_counter.proto --go_out=. --go-grpc_out=. --zrpc_out=.
5. 