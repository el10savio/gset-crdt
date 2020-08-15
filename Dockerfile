FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /gset

WORKDIR /gset

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/gset


FROM scratch

COPY --from=builder /go/bin/gset /go/bin/gset

ENTRYPOINT ["/go/bin/gset"]

EXPOSE 8080