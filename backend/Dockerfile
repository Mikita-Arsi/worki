FROM golang:1.22.5
COPY . /
WORKDIR /
RUN go mod download
#WORKDIR /cmd
RUN go build -o worki main.go
EXPOSE 1323
CMD ["./worki"]