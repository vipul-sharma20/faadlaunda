FROM golang:1.8

RUN mkdir /faad
WORKDIR /faad

ADD . /faad/

RUN go build

EXPOSE 8100

CMD ["go", "run", "faad_server.go", "view.go"]

