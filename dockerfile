FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y libmupdf-dev

RUN go build -o /main .

EXPOSE 5001

CMD ["/main"]
