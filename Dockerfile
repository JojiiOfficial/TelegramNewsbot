FROM golang:1.14.2-alpine3.11 as builder1

# Setting up environment for builder1
ENV GO111MODULE=on
WORKDIR /app/gnews

# install required package(s)
RUN apk --no-cache add ca-certificates git

# Copy dependency list
COPY go.mod .
COPY go.sum .

# Copy files
COPY ./*.go ./

RUN go mod download
# Compile
RUN go build -o main

# Create new stage based on alpine
FROM alpine:latest

#Copy ca certs
COPY --from=builder1 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy compiled binary from builder1
WORKDIR /app
RUN mkdir /app/data/

COPY --from=builder1 /app/gnews/main .

# Run the server
CMD [ "/app/main" ]
