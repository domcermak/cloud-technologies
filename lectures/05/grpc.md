

## Protobuf

1. [Info](https://developers.google.com/protocol-buffers) 
2. Proto2/Proto3
3. Data types
4. Default values
5. Nested types
6. Unknown fields
7. One of

8. [Compilation](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)


## GRPC
1. Definition
    ```protobuf
    syntax = "proto3";
      
    message HelloRequest {
     string greeting = 1;
    }
    
    message HelloResponse {
      string reply = 1;
    }
   
    service HelloService {
      rpc SayHello (HelloRequest) returns (HelloResponse);
      rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
      rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
      rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
    }
   ```
2. [Core concepts](https://grpc.io/docs/what-is-grpc/core-concepts/)
3. Unary/Client stream/Server stream/Bidi stream


## Example project
1. git clone -b v1.43.0 https://github.com/grpc/grpc-go
2. route guide
3. features
