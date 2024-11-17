FROM golang:1.23.2-alpine


WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./
COPY ../config/.env ../../config/

RUN go mod tidy

COPY ../. /app/

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go
 
EXPOSE 8080

CMD ["./app"]

