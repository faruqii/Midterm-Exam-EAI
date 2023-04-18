FROM golang:1.18.0-alpine

RUN apk update && apk add --no-cache git

WORKDIR /go/src/app

COPY . .


ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=postgres
ENV DB_NAME=postgres
ENV DB_SSLMODE=disable

ENV PORT=3000

RUN go mod download && go mod verify

RUN go mod tidy && go mod vendor

RUN chmod +x /go/src/app

ENV WAIT_VERSION 2.7.3
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN go build -o /server cmd/main.go

EXPOSE 3000

CMD ["/server"]