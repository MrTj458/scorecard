FROM golang:1.17-alpine as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v .
RUN go build -v .

FROM alpine:latest

WORKDIR /app

COPY --from=build /go/src/app/scorecard .

CMD ["./scorecard"]
