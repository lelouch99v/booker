FROM golang:latest

ENV SRC_DIR=/go/src/github.com/lelouch99v/booker

ENV GOBIN=/go/bin

WORKDIR $GOBIN

COPY . $SRC_DIR

RUN cd /go/src;
RUN go install github.com/lelouch99v/booker/;

ENTRYPOINT ["./booker"]

EXPOSE 5010
