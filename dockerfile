FROM golang:1.21-alpine

WORKDIR /

COPY src/. .

RUN go mod download

RUN go build -o /cucalen

EXPOSE 5000

CMD [ "/cucalen" ]