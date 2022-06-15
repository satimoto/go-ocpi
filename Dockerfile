FROM golang:alpine AS build-env

RUN apk add -U --no-cache ca-certificates

RUN mkdir /app
WORKDIR /app

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s -w' -o /go/bin/app cmd/ocpi/main.go

FROM scratch

COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build-env /go/bin/app /go/bin/app
EXPOSE 9001 50000
CMD [ "/go/bin/app" ]