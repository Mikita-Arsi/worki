FROM golang:1.22.5

COPY . /go/src/app

WORKDIR /go/src/app/cmd

RUN go build -o worki main.go

EXPOSE 1323

CMD ["./worki"]
