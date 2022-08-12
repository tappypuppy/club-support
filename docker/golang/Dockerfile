FROM golang:alpine
RUN apk update
RUN apk add alpine-sdk
COPY ./go-app /app/go
WORKDIR /app/go
RUN apk add --no-cache git 
RUN apk add bash
# added
RUN ls
RUN go mod download
CMD go run main.go