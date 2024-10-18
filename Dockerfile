FROM golang:1.23.0-alpine3.20

RUN apk update && apk upgrade
RUN apk add --no-cache  \
    build-base  \
    make

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

RUN make build

CMD ["make", "watch"]