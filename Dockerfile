FROM golang:1.24.2 as builder
WORKDIR /app
COPY . .
RUN go build -o server .

FROM alpine
COPY --from=builder /app/server /server
CMD ["/server"]
