FROM golang:1.16-alpine

RUN mkdir -p /app/

WORKDIR /app/

COPY . /app/

RUN go install github.com/acoshift/goreload@master

ENTRYPOINT ["goreload", "main.go"]

EXPOSE 3000