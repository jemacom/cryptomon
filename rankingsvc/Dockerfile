FROM golang as builder
WORKDIR /go/src/github.com/jemacom/cryptomon/rankingsvc

RUN go get -v github.com/lucazulian/cryptocomparego

COPY ["cryptocompare.go", "main.go", "./"]

RUN CGO_ENABLED=0 GOOS=linux go build -o rankingsvc .



FROM alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/jemacom/cryptomon/rankingsvc/rankingsvc /

CMD [ "/rankingsvc", "-port=8082", "2>/dev/null &"]

EXPOSE 8080
