FROM golang:latest

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/server
RUN git submodule update --init --recursive

EXPOSE 80
EXPOSE 443
STOPSIGNAL SIGINT
CMD ["app"]
