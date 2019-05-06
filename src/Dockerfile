FROM golang:1.12.0-alpine

# install git
RUN apk update && \
    apk upgrade && \
    apk add git

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

# copy core and lib
COPY xchain/ /go/src/xchain-contract-go/src/xchain