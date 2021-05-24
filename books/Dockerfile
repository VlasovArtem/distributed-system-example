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
RUN go build -o /out/books cmd/*.go

FROM alpine:3.13.3 AS bin
ENV HTTP_PORT=8080
ENV RPC_TCPPORT=50052
COPY --from=build /out/books /
EXPOSE ${HTTP_PORT}
ENTRYPOINT ["/books"]