FROM golang:alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/main /app
COPY --from=builder /build/*.txt /app

ENTRYPOINT [ "/app/main" ]
CMD [ "" ]