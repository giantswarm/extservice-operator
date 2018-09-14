FROM alpine:3.8

RUN apk add --no-cache ca-certificates

ADD ./extservice-operator /extservice-operator

ENTRYPOINT ["/extservice-operator"]
