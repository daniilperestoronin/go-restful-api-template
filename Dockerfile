FROM golang as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o go-restful-api-template .

FROM alpine:latest
ENV DB_DRIVER="postgres"\
    DB_DATA_SOURCE="host=172.17.0.1 port= 5432 user=postgres password=postgres dbname=records sslmode=disable"
COPY --from=builder /build/go-restful-api-template /app/
WORKDIR /app
CMD ["./go-restful-api-template"]