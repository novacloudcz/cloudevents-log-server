FROM golang as builder
WORKDIR /go/src/github.com/graphql-services/memberships
COPY . .
RUN go get ./... 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /tmp/app main.go

FROM alpine:3.5

WORKDIR /app

COPY --from=builder /tmp/app /usr/local/bin/app

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /usr/local/bin/app

ENTRYPOINT []
CMD [ "/usr/local/bin/app"]