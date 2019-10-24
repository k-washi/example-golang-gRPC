# example-golang-gRPC

Rest vs gRPC: https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html

## SetUp

---

### Vscode setup

+ Install  vscode-proto3 & Clang-Format extensions

+ Install Clang-Format(if you want to format your files automatically)

```bash
brew install clang-format #MacOSX

#other http://www.codepool.biz/vscode-format-c-code-windows-linux.html
```

### Protoc Setup

+ install protoc

```bash
brew install protobuf #MacOSX
```

linux: https://github.com/google/protobuf/releases

```bash
#example
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip
# Unzip
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3
# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/
# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/
# Optional: change owner
sudo chown [user] /usr/local/bin/protoc
sudo chown -R [user] /usr/local/include/google
```

### golang Setup

```bash
go get -u google.golang.org/grpc
go get -d -u github.com/golang/protobuf/protoc-gen-go
```

### test

```bash
protoc greet/greetpb/greet.proto  --go_out=plugins=grpc:.

#server setup
go run greet/greet_server/server.go 

#client setup
go run greet/greet_client/client.go
```

## Unary

---

Unary RPC calls wil be the most common for APIs.
The client will send one message to the server and will receive on response from server.

example: "unary"

