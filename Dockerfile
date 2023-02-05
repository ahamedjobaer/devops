FROM golang:1.19-alpine as base

FROM base as builder
WORKDIR /app
COPY . .
RUN go build -o ./build/server

FROM base as prod
WORKDIR /app
COPY --from=builder /app/build/server ./build/server

EXPOSE 8080

CMD ["/app/build/server"]
