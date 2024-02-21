FROM golang
COPY . .
RUN apt-get update && apt-get install unzip
RUN go build login.go