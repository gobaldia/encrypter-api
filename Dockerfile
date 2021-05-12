FROM golang:1.15.7-buster

ENV PORT=3000

COPY . /app

WORKDIR /app

RUN go build

EXPOSE $PORT

ENTRYPOINT [ "./encrypter-api" ]