FROM golang:1.14 AS builder
WORKDIR /go/src/restapi
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o restapi main.go

FROM scratch
COPY --from=builder /go/src/restapi/restapi .
EXPOSE 8080:8080
ENTRYPOINT [ "./restapi" ]
CMD [ "" ]