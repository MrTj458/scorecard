FROM node:17-alpine as client-build

WORKDIR /app

COPY ./client .

RUN npm install

RUN npm run build

FROM golang:1.17-alpine as server-build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v .
RUN go build -v .

FROM alpine:3

WORKDIR /app

COPY --from=server-build /go/src/app/scorecard .
COPY --from=client-build /app/build ./static

CMD ["./scorecard"]
