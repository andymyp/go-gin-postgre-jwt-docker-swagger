FROM golang:1.22.5

WORKDIR /go/src/github.com/andymyp/go-gin-postgre-jwt-docker-swagger

COPY . .

RUN go get -v

RUN go build -o main

CMD [ "./main" ]