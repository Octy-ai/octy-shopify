#MULTISTAGE DOCKERFILE
FROM golang:1.16.6 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux go build -o octy-shopify -ldflags '-w -extldflags "-static"' ./cmd

FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 8080
CMD ["./octy-shopify"]