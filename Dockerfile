FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o catopia main.go
RUN go build -o migrate cmd/migration.go

FROM debian:latest
ENV TZ="Asia/Bangkok"
WORKDIR /root/
COPY --from=0 /app/catopia .
COPY --from=0 /app/migrate .
COPY --from=0 /app/others ./others
EXPOSE 8080
CMD ["./migrate","./catopia"]