FROM golang:1.14

LABEL base.name="showbasetest"

WORKDIR /app/backend

COPY . .

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["./main"]


