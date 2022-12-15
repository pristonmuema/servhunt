FROM golang:1.18 AS build-env

ENV GO_ENV=dev
ENV TZ=Africa/Nairobi
RUN mkdir -p /go/src/servhunt
WORKDIR /go/src/servhunt
COPY go.mod go.sum ./
RUN go mod download

COPY ../servhunte .

RUN  go build  -o servhunt .

WORKDIR /root/

COPY --from=build-env /go/src/servhunt .

EXPOSE 9094

CMD ["./servhunt"]
