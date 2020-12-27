FROM golang:1.12.0-alpine3.9
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go build -o api .
CMD ["/app/api"]