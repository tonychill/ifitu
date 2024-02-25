#!/bin/bash
set -e

function golang() {
  rm -rf ./pb/go
  rm -rf ./pb/rust
  rm -rf ./pb/java
  mkdir -m 777 ./pb/go
  mkdir -m 777 ./pb/rust
  # curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v21.6/protoc-21.6-linux-aarch_64.zip" && \
  #   unzip "protoc-21.6-linux-aarch_64.zip" -d $HOME/.local
  # protoc -I=$HOME/.local/include -I=./protos ./protos/*.proto \
  protoc -I=$HOME/.local/include -I=./protos ./protos/*.proto \
    --doc_out=./pb/go --doc_opt=markdown,README.md \
    --go_out=./pb/go --go_opt=paths=import \
    --go-grpc_out=./pb/go --go-grpc_opt=paths=import \
    --rust_out=./pb/rust 
    # --ts_out=/Users/tonyhill/wuk/juvae/software/projects/juvae/clients/web_guests/src/protos \
    # ./protos/*.proto
  mv ./pb/go/github.com/tonychill/ifitu/apis/pb/go/* ./pb/go
  rm -r ./pb/go/github.com
  # cd ./pb/go && $(go mod init github.com/tonychill/ifitu/apis/pb/go)  && \
    # $(go mod tidy)
  echo "done doing the thang 😜"
}
golang
