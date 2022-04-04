FROM golang:alpine
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]