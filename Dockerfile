FROM golang:1.23.2-alpine AS builder

WORKDIR /build

COPY . .

# RUN go mod init gosql
RUN go mod tidy && go mod vendor
RUN cd /build/app && go build -o main
# RUN cd /app && go build -o main

FROM alpine:3.21

COPY --from=builder /build/app/main /app/main
COPY --from=builder /build/app/.env /app/.env
EXPOSE 8080

WORKDIR /app
CMD ["./main"]
# --command="./main"

## Running container
# docker run --publish 8080:8080 my-go-app ./main