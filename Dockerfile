FROM golang:1.21.6-bookworm

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /go-docker

EXPOSE 8123

CMD [ "/go-docker"]