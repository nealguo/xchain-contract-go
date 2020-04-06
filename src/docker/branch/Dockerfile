FROM xchain-contract-goenv:base

# copy tls and core
COPY data/* /opt/app/
COPY xchain/ /go/src/xchain-contract-go/src/xchain
WORKDIR /go/src/xchain-contract-go/src/xchain

# build then copy
RUN go build -p 1 -o usercontract
RUN cp -f /go/src/xchain-contract-go/src/xchain/usercontract /opt/app

# run the server
WORKDIR /opt/app
ENTRYPOINT ["./usercontract"]

# NOTE: this Dockerfile will build xchain-contract-goenv:3.8