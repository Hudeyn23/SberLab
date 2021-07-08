FROM alpine:latest

WORKDIR /opt/sbapi

EXPOSE 8080

ADD ./target/sbapi sbapi

ENTRYPOINT ["/opt/sbapi/sbapi"]
