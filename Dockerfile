FROM alpine

COPY gopath/bin/go-pipeline /go/bin/go-pipeline

ENTRYPOINT /go/bin/go-pipeline