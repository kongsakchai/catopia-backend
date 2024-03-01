FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o catopia ./src/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /root/
COPY --from=0 /app/catopia .
EXPOSE 8080
CMD ["./catopia"]
