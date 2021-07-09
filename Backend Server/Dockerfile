FROM golang:1.16.5-buster

WORKDIR /back
COPY ./ .
EXPOSE 9999

RUN go build

CMD [ "./backend" ]
