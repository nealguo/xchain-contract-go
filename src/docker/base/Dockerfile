FROM golang:1.12.0-alpine

# install libs
RUN apk update
RUN apk upgrade
RUN apk add git
RUN apk add gcc
RUN apk add musl-dev

# install grpc
RUN mkdir -p $GOPATH/src/google.golang.org
RUN cd $GOPATH/src/google.golang.org
RUN git clone https://github.com/grpc/grpc-go.git grpc
RUN git clone https://github.com/google/go-genproto.git genproto
RUN mkdir -p $GOPATH/src/golang.org/x
RUN cd $GOPATH/src/golang.org/x
RUN git clone https://github.com/golang/net.git
RUN git clone https://github.com/golang/text.git
RUN git clone https://github.com/golang/sys.git
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/ptypes
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go install google.golang.org/grpc

RUN echo "succeeded to install libs and grpc"

# NOTE: this Dockerfile will build xchain-contract-goenv:base