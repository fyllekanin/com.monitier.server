FROM golang:1.19

WORKDIR /app
RUN apt update -y && apt upgrade -y
RUN apt install build-essential -y

COPY ./ ./
RUN go mod download
RUN go build

EXPOSE 8080

CMD ["./com.monitier.server"]