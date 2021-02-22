FROM golang:1.15.8-alpine as build

ARG CLEARDB_DATABASE_URL
ENV CLEARDB_DATABASE_URL $CLEARDB_DATABASE_URL

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o go-money-api ./cmd/go-money-api/main.go

FROM alpine:3.13.1

COPY --from=build /go/src/app/go-money-api .
COPY --from=build /go/src/app/scripts scripts

ENTRYPOINT [ "./go-money-api" ]