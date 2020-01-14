FROM golang:latest as builder
ENV SRC_DIR=/go/src/github.com/lelouch99v/booker
ENV GOBIN=/go/bin
WORKDIR $GOBIN
COPY . $SRC_DIR
RUN cd /go/src && go install github.com/lelouch99v/booker/

# runtime image
FROM alpine:latest
COPY --from=builder /go/bin/booker /booker
EXPOSE 5010
ENTRYPOINT ["/booker"]
