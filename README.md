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
```sql
-- 用户信息表
CREATE TABLE user
(
    user_id  bigint AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL DEFAULT '',
    PRIMARY KEY (user_id)
);

-- 文章信息表
CREATE TABLE article
(
    article_id bigint AUTO_INCREMENT,
    title      VARCHAR(255) NOT NULL DEFAULT '',
    content    TEXT         NOT NULL DEFAULT '',
    author_id  bigint       NOT NULL,
    PRIMARY KEY (article_id)
);

-- 点赞信息表
CREATE TABLE likes
(
    like_id    bigint AUTO_INCREMENT,
    user_id    bigint NOT NULL,
    article_id bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (article_id)
);

-- 收藏信息表
CREATE TABLE favorites
(
    favorite_id bigint AUTO_INCREMENT,
    user_id     bigint NOT NULL,
    article_id  bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    PRIMARY KEY (favorite_id)
);
```
2. cd model/mysql && goctl model mysql ddl --src ddl.sql --dir . -c
3. define .proto file
```protobuf
syntax = "proto3";

package go_counter;
option go_package="./go_counter";

message PingRequest {
  string ping = 1;
}

message PingResponse {
  string pong = 1;
}


message LikeRequest {
  string user_id = 1;
  string content_id = 2;
}

message LikeResponse {
  bool success = 1;
}

message FavoriteRequest {
  string user_id = 1;
  string content_id = 2;
}

message FavoriteResponse {
  bool success = 1;
}

message ViewRequest {
  string user_id = 1;
  string content_id = 2;
}

message ViewResponse {
  bool success = 1;
}


service Go_counter {
  rpc Ping(PingRequest) returns(PingResponse);

  // 点赞请求
  rpc Like(LikeRequest) returns (LikeResponse);

  // 收藏请求
  rpc Favorite(FavoriteRequest) returns (FavoriteResponse);

  // 浏览请求
  rpc View(ViewRequest) returns (ViewResponse);
}

```
4. goctl rpc protoc go_counter.proto --go_out=. --go-grpc_out=. --zrpc_out=.
5. 