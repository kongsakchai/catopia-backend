FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o catopia main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat tzdata
ENV TZ="Asia/Bangkok"
WORKDIR /root/
COPY --from=0 /app/catopia .
EXPOSE 8080
CMD ["./catopia"]
