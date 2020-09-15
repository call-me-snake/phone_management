FROM golang:1.12.0-alpine3.9 AS builder
WORKDIR /go/src/github.com/call-me-snake/phone_management
COPY . .
RUN go install ./...

FROM jwilder/dockerize AS production
COPY --from=builder /go/bin/cmd ./app

#docker build -t phone_management_img .
#docker run -it --name phone_management phone_management_img /bin/sh
#docker stop phone_management
#docker rm phone_management

