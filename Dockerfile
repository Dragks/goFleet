FROM golang:1.19-alpine

EXPOSE 8080

RUN apk update \
  && apk add --no-cache \
    mysql-client \
    build-base
  
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN /bin/chmod +x /usr/local/bin/entrypoint.sh

RUN go build cmd/main.go
RUN mv main /usr/local/bin/

CMD ["main"]
ENTRYPOINT ["entrypoint.sh"]
