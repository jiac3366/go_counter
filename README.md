# counter service


### 
1. 写性能差 --> 异步写
    1.1 在计数服务端聚合写  (缺点: 中途宕机会丢数据，少记)
        适合少记不敏感的数据 (例如: 浏览量) 
        例如：1 分钟写一次 redis，1 小时写一次 MySQL ？

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
    user_id  bigint NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL DEFAULT '',
    PRIMARY KEY (user_id)
);

-- 文章信息表
CREATE TABLE article
(
    article_id bigint NOT NULL AUTO_INCREMENT,
    title      VARCHAR(255) NOT NULL DEFAULT '',
    content    TEXT         NOT NULL DEFAULT '',
    author_id  bigint       NOT NULL,
    PRIMARY KEY (article_id)
);

-- 点赞记录表
CREATE TABLE likes_record
(
    like_id    bigint NOT NULL AUTO_INCREMENT,
    user_id    bigint NOT NULL,
    article_id bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    UNIQUE KEY idx_uniq_user_id_article_id (user_id, article_id),
    PRIMARY KEY (like_id)
);


-- 收藏记录表
CREATE TABLE favorites_record
(
    favorite_id bigint NOT NULL AUTO_INCREMENT,
    user_id     bigint NOT NULL,
    article_id  bigint NOT NULL,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    UNIQUE KEY idx_uniq_user_id_article_id (user_id, article_id),
    PRIMARY KEY (favorite_id)
);

-- 计数表
CREATE TABLE count_meta
(
    id  bigint NOT NULL auto_increment,
    business_id bigint NOT NULL,  -- same as article_id
    count bigint NOT NULL,
    types VARCHAR(32) NOT NULL,
    UNIQUE KEY idx_uniq_business_id (business_id, types), -- article_id_1, "FAVORITE" 代表 id 为 1 的文章点赞数
    PRIMARY KEY (id)
);

```
2. cd model && goctl model mysql ddl --src ddl.sql --dir ./mysql -c

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


message LikeNumRequest {
  int64 article_id = 1;
}

message LikeNumResponse {
  bool count = 1;
}

message LikeRecordRequest {
  int64 user_id = 1;
  int64 article_id = 2;
}

message LikeRecordResponse {
  bool success = 1;
}

// view the count of favorite
message FavoriteNumRequest {
  int64 article_id = 1;
}

message FavoriteNumResponse {
  int64 count = 1;
}

// favorite record operations
message FavoriteRecordRequest {
  int64 user_id = 1;
  int64 article_id = 2;
}

message FavoriteRecordResponse {
  bool success = 1;
}



message ViewNumRequest {
  int64 article_id = 1;
}

message ViewNumResponse {
  bool success = 1;
}


message ViewRecordRequest {
  int64 user_id = 1;
  int64 article_id = 2;
}

message ViewRecordResponse {
  bool success = 1;
}


service Go_counter {
  rpc Ping(PingRequest) returns(PingResponse);

  // 点赞请求
  rpc LikeNum(LikeNumRequest) returns (LikeNumResponse);
  rpc LikeRecord(LikeRecordRequest) returns (LikeRecordResponse);

  // 收藏请求
  rpc FavoriteNum(FavoriteNumRequest) returns (FavoriteNumResponse);
  rpc FavoriteRecord(FavoriteRecordRequest) returns (FavoriteRecordResponse);

  // 浏览请求
  rpc ViewNum(ViewNumRequest) returns (ViewNumResponse);
  rpc ViewRecord(ViewRecordRequest) returns (ViewRecordResponse);
}
```
4. goctl rpc protoc go_counter.proto --go_out=. --go-grpc_out=. --zrpc_out=.