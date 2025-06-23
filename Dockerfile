FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o mailing-service
EXPOSE 9000
CMD ["./mailing-service"]
