FROM golang:1.16

WORKDIR /app/
COPY ./ ./

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 8080
