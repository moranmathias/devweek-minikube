FROM  golang:alpine as builder
COPY "." "$GOPATH/src/github/moranmathias/devweek"
WORKDIR $GOPATH/src/github.com/moranmathias/devweek
#Build the binary
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./cmd/echosrv/main.go

FROM alpine
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]