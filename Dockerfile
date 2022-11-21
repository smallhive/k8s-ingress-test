FROM alpine:3.16.3

ADD ./echo-service .
CMD ["/echo-service"]
