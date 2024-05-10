FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o catopia main.go
RUN go build -o migrate cmd/migrate.go

FROM debian:stable-slim
ENV TZ="Asia/Bangkok"
WORKDIR /root/
COPY --from=0 /app/catopia .
COPY --from=0 /app/migrate .
COPY --from=0 /app/start.sh .
EXPOSE 8080
CMD ["sh","start.sh"]