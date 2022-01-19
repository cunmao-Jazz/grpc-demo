# 如何编译 proto文件

```
# 进入到当前probuf定义的目录
$ cd protocol

# 指定protobuf 文件的搜索位置为当前目录 I=.
$ protoc -I=. -I=/usr/local/include --go_out=. --go-grpc_out=. --go-grpc_opt=module="github.com/cunmao-Jazz/grpc-demo/protocol" --go_opt=module="github.com/cunmao-Jazz/grpc-demo/protocol" hello.proto
