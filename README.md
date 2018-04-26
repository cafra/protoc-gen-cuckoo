# protoc-gen-cuckoo

**cuckoo** RPC support many codecs, if you use **Protocol Buffers** codec, you can generate contract codes from [Protocol Buffers](https://developers.google.com/protocol-buffers/docs/proto3) service definition files with **protoc-gen-cuckoo**.

## Install

Use `go get` to install the code generator:

```bash
go install github.com/carolove/protoc-gen-cuckoo
```

You will also need:

* [protoc](https://github.com/golang/protobuf), the protobuf compiler. You need version 3+.
* [github.com/golang/protobuf/protoc-gen-go](https://github.com/golang/protobuf/), the Go protobuf generator plugin. Get this with `go get`.

## Usage

Just like **grpc**:

```bash
protoc --go_out=. --cuckoo_out=. hello.proto
```

Service interfaces and client proxies were generated into a separate file `[name].cuckoo.go`:

```
hello.cuckoo.go
hello.pb.go
hello.proto
```