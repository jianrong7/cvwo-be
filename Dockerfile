FROM golang:alpine as build

ENV GIN_MODE=release

WORKDIR /go/src/cvwo-be

COPY . ./

RUN go build

EXPOSE 3000

ENTRYPOINT [ "./cvwo-be" ]