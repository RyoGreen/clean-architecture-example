FROM golang:1.21 AS build

COPY . /app/
WORKDIR /app/

RUN go get github.com/lib/pq
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/app/

FROM golang:1.21

COPY --from=build /app/app /app/app
# COPY --from=build /app/dist /app/dist
# COPY --from=build /app/tpls /app/tpls
# COPY --from=build /app/config.json /app/config/json

WORKDIR /app

CMD ["/app/app"]
