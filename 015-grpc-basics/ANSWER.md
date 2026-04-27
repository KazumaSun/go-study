# 015 正解例（課題別）

## 課題1: proto定義（unary）

```proto
syntax = "proto3";

package user.v1;

option go_package = "example.com/userpb;userpb";

service UserService {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
}

message GetUserRequest {
  string id = 1;
}

message GetUserResponse {
  string id = 1;
  string name = 2;
}
```

## 課題2: server streaming（proto）

```proto
rpc ListUsers(ListUsersRequest) returns (stream User);

message ListUsersRequest {}

message User {
  string id = 1;
  string name = 2;
}
```

## 課題3: interceptor最小例

```go
func authUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md.Get("authorization")) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing authorization")
	}
	return handler(ctx, req)
}
```
