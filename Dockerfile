from golang:1.22-alpine3.20

workdir /go/src/app
copy . .
run go mod vendor
run go build -o /usr/local/bin/app

from alpine:3.20
copy --from=0 /usr/local/bin/app /usr/local/bin/app
expose 8989
entrypoint ["/usr/local/bin/app"]
