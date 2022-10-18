##
## Build Stage
##

FROM golang:1.19 AS builder

WORKDIR /workspace

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /deviceplugin

##
## Deploy Stage
##

FROM alpine:3.15

WORKDIR /

COPY --from=builder /deviceplugin /deviceplugin

USER 1001

ENTRYPOINT ["/deviceplugin"]
