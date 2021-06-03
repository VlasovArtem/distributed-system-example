FROM golang:1.16.2 AS base
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

FROM base AS build
RUN go build -o /out/authors cmd/*.go

FROM alpine:3.13.3 AS bin
ENV RPC_TCPPORT=50051
COPY --from=build /out/authors /
EXPOSE ${HTTP_PORT}
ENTRYPOINT ["/authors"]