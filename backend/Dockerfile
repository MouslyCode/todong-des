# Stage 1: Build
FROM golang:1.26.4-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy 
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]