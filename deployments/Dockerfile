FROM golang:1.15.8-alpine as build

ARG DATABASE_URL
ENV DATABASE_URL $DATABASE_URL

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o go-money-api ./cmd/go-money-api/main.go

FROM alpine:3.13.1

COPY --from=build /go/src/app/go-money-api .

ENTRYPOINT [ "./go-money-api" ]