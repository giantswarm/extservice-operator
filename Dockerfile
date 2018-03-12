FROM alpine:3.5

RUN apk add --no-cache ca-certificates

ADD ./extservice-operator /extservice-operator

ENTRYPOINT ["/extservice-operator"]
