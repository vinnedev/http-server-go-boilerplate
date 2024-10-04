FROM golang:1.22-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o /app/main ./cmd
EXPOSE 80

CMD ["/app/main"]