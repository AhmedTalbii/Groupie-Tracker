FROM golang:1.22-alpine AS builder

WORKDIR temp

COPY . .

RUN go build -o main main.go

FROM alpine:latest

WORKDIR gruopieTrucker

COPY --from=builder /temp/main .
COPY --from=builder /temp/views .
COPY --from=builder /temp/biblio/pages .

CMD ["./main"]