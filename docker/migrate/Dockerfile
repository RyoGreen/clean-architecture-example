FROM golang:1.21 AS build

COPY . /app/
WORKDIR /app/

RUN go get ./... \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/migrate/

FROM golang:1.21

COPY --from=build /app/migrate /app/migrate
COPY --from=build /app/config.json /app/config.json

WORKDIR /app

CMD ["/app/migrate"]
