# develop stage
FROM golang:alpine as develop-stage
RUN apk add --no-cache git
RUN mkdir /app
ADD . /app
WORKDIR /app

# build stage
FROM develop-stage as build-stage
RUN go build -o cocopasty-backend .

# production stage
FROM golang:alpine as production-stage
COPY --from=build-stage /app/cocopasty-backend /app/cocopasty-backend
CMD ["/app/cocopasty-backend"]