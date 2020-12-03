FROM golang:1.15 as builder
WORKDIR /app
ADD go.mod go.sum ./
RUN go mod download
ADD main.go ./
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=builder /app/client-go-example /app/
ENTRYPOINT ["/app/client-go-example"]
