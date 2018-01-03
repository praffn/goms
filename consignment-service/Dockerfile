FROM golang:1.9.0 as builder

# set workdir to our current service in the gopath
WORKDIR /go/src/github.com/praffn/goms/consignment-service

# copy the code to our workdir
COPY . .

# get godep (dependency management tool)
RUN go get -u github.com/golang/dep/cmd/dep

# create dep project and run 'ensure' to pull all dependencies
RUN dep init && dep ensure

# build the binary with a few flags which will allow us
# to run this binary in Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# now start new build process with this image
FROM alpine:latest

# security
RUN apk --no-cache add ca-certificates

# create workdir for our app
RUN mkdir /app
WORKDIR /app

# pull binary from builder
COPY --from=builder /go/src/github.com/praffn/goms/consignment-service/consignment-service .

# and run binary
CMD ["./consignment-service"]