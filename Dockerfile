FROM golang:1.18.2 AS build
WORKDIR /server

COPY main.go .
COPY go.mod .
COPY go.sum .

RUN go mod download

# compile server
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app .

FROM scratch
WORKDIR /root/
COPY --from=build /server/app .
CMD ["./app"]
