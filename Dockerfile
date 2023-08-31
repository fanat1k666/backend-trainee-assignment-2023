FROM golang:alpine

WORKDIR /build
COPY . .
RUN go build -o ab cmd/ab/main.go
EXPOSE 8000
CMD "/build/ab"