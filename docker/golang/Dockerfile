FROM golang:alpine
RUN apk update
RUN apk add alpine-sdk
COPY ./go-app /app/go
WORKDIR /app/go
RUN apk add --no-cache git 
RUN go mod init example
RUN go get github.com/gin-gonic/gin
CMD go run main.go
# RUN go build -o main .
# CMD [ "/app/go/main" ]