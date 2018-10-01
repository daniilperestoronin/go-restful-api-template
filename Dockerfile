FROM golang as builder
ADD . ./app/
RUN go get -d -v
RUN go build

FROM alpine:latest
COPY --from=builder /app/go-restful-api-template .
CMD ["./go-restful-api-template"]