FROM golang:1.20-alpine as builder
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/github.com/arieffian/go-boilerplate
COPY . .
#RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /cmd/app/main
# need to load timezone info and copy it over manually
RUN apk --no-cache add tzdata
COPY ./configs /configs
COPY ./migrations /migrations

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /main /main
COPY --from=builder /migrations /migrations
COPY --from=builder /.env /.env
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

ARG env
ENV ENVIRONMENT $env
USER appuser
EXPOSE 8080 8081
ENTRYPOINT ["/main"]
