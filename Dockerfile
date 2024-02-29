FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o catopia ./src/main.go

FROM alpine:latest
ARG PORT=8080
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /root/
COPY --from=0 /app/catopia .
EXPOSE $PORT
CMD ["./catopia"]
