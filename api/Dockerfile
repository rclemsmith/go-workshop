FROM golang:1.20-alpine AS BuildStage
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-demo

EXPOSE 8080

CMD [ "/go-demo" ]

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /go-demo /go-demo

EXPOSE 8080 

RUN chmod +x /go-demo

ENTRYPOINT ["/go-demo"]