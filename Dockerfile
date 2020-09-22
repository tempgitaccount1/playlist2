FROM golang
ADD . /go/src/playlist
RUN go get github.com/golang/protobuf/proto
RUN go get google.golang.org/grpc
RUN go install playlist
ENTRYPOINT ["/go/bin/playlist"]
EXPOSE 8080

