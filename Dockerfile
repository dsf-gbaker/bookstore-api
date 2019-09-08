FROM golang:1.12.9 AS builder
WORKDIR /go/src/restapi
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o restapi main.go

FROM scratch
COPY --from=builder /go/src/restapi/restapi .
# Use the below command if you've built the program
# using the docker-compose which writes out the file
# to ./bin/
#COPY ./bin/restapi .
EXPOSE 8080:8080
ENTRYPOINT [ "./restapi" ]
CMD [ "" ]