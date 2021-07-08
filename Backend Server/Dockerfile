FROM alpine:latest

WORKDIR /opt/sbapi

EXPOSE 9999

ADD ./target/sbapi sbapi

ENTRYPOINT ["/opt/sbapi/sbapi"]
